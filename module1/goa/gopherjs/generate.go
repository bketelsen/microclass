//go:generate goagen -d github.com/bketelsen/microclass/module1/goa/gopherjs/design app
//go:generate goagen -d github.com/bketelsen/microclass/module1/goa/gopherjs/design main
//go:generate goagen -d github.com/bketelsen/microclass/module1/goa/gopherjs/design client
//go:generate goagen -d github.com/bketelsen/microclass/module1/goa/gopherjs/design swagger
//go:generate gopherjs build github.com/bketelsen/microclass/module1/goa/gopherjs/public -o public/website.js
//go:generate mkdir -p images
package main
