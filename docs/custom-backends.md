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

Install the backend during process startup. Do not replace it while an application
is running. Dispatch input, painting and service callbacks on the host's serial
event loop, and signal `OnClosed` exactly once for each window.

A host that implements `app.HostedBackend` and returns true from
`ExternallyDriven` owns the event loop: `RunCanvas` and `RunApplication` register
the application and return. Resource cleanup must follow the window's lifecycle;
a `defer` immediately after `RunCanvas` no longer means "after the window closes."

Backend contract tests run without native windows:

```sh
go test -tags viva_custom_backend ./app
go build -tags viva_custom_backend ./...
```

These tests cover production painter frames and service ownership. Actual GPU,
input-method, window and service behavior must also be checked in each adapter's
host environment. A public contract does not imply every adapter implements all
optional capabilities.
