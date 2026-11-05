# Custom native backends

An external host can install `app.Backend` with `app.UseBackend` before starting
an application. The backend receives the same painter frames and input contracts
as the built-in platforms. Controls, layout and application code keep using the
normal Viva APIs; platform implementation belongs to the external adapter.

Build with `-tags viva_custom_backend` to exclude the built-in desktop and browser
backends. The host must supply a text engine, window/event handling and a renderer.
`app.BackendConfig` and `app.BackendFrame`, together with the other `Backend*`
aliases, expose the existing contracts without importing internal packages.
Frames are borrowed: consume them before invoking another framework callback.

`Backend.Services` supplies optional platform operations. Once a backend is
selected, absent services report `app.ErrUnsupported` or the documented empty
result; they do not call a built-in platform service. Shared text editing,
clipboard and settings helpers preserve the framework's normal behavior.
`app.OffscreenBackend` is a separate optional capability for headless rendering.

An adapter can implement `OffscreenBackend.NewOffscreen` by returning
`app.NewRasterOffscreen(config, rasterize)`. This pure-Go helper renders the
scene without a native window or GPU. Width and height are positive logical
pixel counts; scale is from 1 to 64, with each device dimension rounded up and
limited to 16384 pixels. It draws shapes, bitmap labels, text labels, the active
editor's text, and layered holes for native views.

The optional `app.BackendLabelRasterizer` callback receives a `BackendTextLabel`,
the scale and the required device width and height. Return tightly packed,
premultiplied BGRA pixels, top row first, or an error. The returned buffer belongs
to the renderer; do not reuse it. Bitmap labels bypass the callback and are
copied or downsampled as needed. A nil callback omits text. Secure editor text
is masked with one `•` per rune before the callback; hidden native overlays are
omitted. `Set` consumes borrowed frames and respects `TextInputUpdate.Changed`
and the nil-versus-empty `NativeViews` update semantics.

`Image` returns a fresh premultiplied `image.RGBA`. A clear color with alpha at
least 1 forces opaque output, including native-view holes; otherwise the canvas
alpha is preserved. `Close` releases buffers and subsequent `Set` or `Image`
calls fail. Use each target from one goroutine at a time.

`BackendServices.SetAutostart` and `AutostartEnabled` own login-start changes and
queries for a selected backend. The public `app.SetAutostart` and
`app.AutostartEnabled` APIs validate entries and IDs before dispatching. Missing
callbacks return `app.ErrUnsupported`; no desktop login files or registry entries
are accessed through these services in a custom-backend build.

`BackendServices.DrawsWindowControls` returns true when Viva should paint and
handle minimize, maximize and close buttons in a custom titlebar. A missing
callback returns false. `BackendServices.TitlebarInsets` supplies the leading
and trailing reserved widths in logical pixels for `app.TitlebarInsets`; a
missing callback returns zero for both. A backend that enables Viva's buttons
normally reserves `3 * app.WindowControlWidth` at the trailing end. With no
backend selected, the built-in platform behavior applies.

Install the backend during process startup. Do not replace it while an application
is running. Dispatch input, painting and service callbacks on the host's serial
event loop, and signal `OnClosed` exactly once for each window.

A host that implements `app.HostedBackend` and returns true from
`ExternallyDriven` owns the event loop: `RunCanvas` and `RunApplication` register
the application and return. Resource cleanup must follow the window's lifecycle;
a `defer` immediately after `RunCanvas` no longer means "after the window closes."

Backend contract tests run without native windows:

```sh
go test -tags viva_custom_backend ./app ./app/singleinstance ./headless/...
go test ./app ./internal/raster
go build -tags viva_custom_backend ./...
```

These tests cover production painter frames and service ownership. Actual GPU,
input-method, window and service behavior must also be checked in each adapter's
host environment. A public contract does not imply every adapter implements all
optional capabilities.
