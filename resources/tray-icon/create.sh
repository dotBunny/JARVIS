go get github.com/cratonica/2goarray

tiffutil -cathidpicheck status/icon-error.png status/icon-error@2x.png -out output/icon-error.tiff
tiffutil -cathidpicheck status/icon-ready.png status/icon-ready@2x.png -out output/icon-ready.tiff
tiffutil -cathidpicheck status/icon-warning.png status/icon-warning@2x.png -out output/icon-warning.tiff

/Users/reapazor/go/bin/2goarray TrayIconError main < ./output/icon-error.tiff |  grep -v package >> ./generated/icon-error-unix.go
/Users/reapazor/go/bin/2goarray TrayIconWarning main < ./output/icon-warning.tiff |  grep -v package >> ./generated/icon-warning-unix.go
/Users/reapazor/go/bin/2goarray TrayIconReady main < ./output/icon-ready.tiff |  grep -v package >> ./generated/icon-ready-unix.go

/Users/reapazor/go/bin/2goarray TrayIconError main < ./status/icon-error.ico |  grep -v package >> ./generated/icon-error-win.go
/Users/reapazor/go/bin/2goarray TrayIconWarning main < ./status/icon-warning.ico |  grep -v package >> ./generated/icon-warning-win.go
/Users/reapazor/go/bin/2goarray TrayIconReady main < ./status/icon-ready.ico |  grep -v package >> ./generated/icon-ready-win.go
