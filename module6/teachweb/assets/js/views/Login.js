var m = require("mithril")
var NavMenu = require("./NavMenu")

module.exports = {
    view: function() {
        return (
            <body>
            <NavMenu />
  <div class="wrapper wrapper-full-page">
        <div class="full-page login-page" data-color="" data-image="/assets/img/background/background-2.jpg">
            <div class="content">
                <div class="container">
                    <div class="row">
                        <div class="col-md-4 col-sm-6 col-md-offset-4 col-sm-offset-3">
                            <form method="#" action="#">
                                <div class="card" data-background="color" data-color="blue">
                                    <div class="header">
                                        <h3 class="title">Login</h3>
                                    </div>
                                    <div class="content">
                                        <div class="form-group">
                                            <label>Email address</label>
                                            <input type="email" placeholder="Enter email" class="form-control input-no-border"></input>
                                        </div>
                                        <div class="form-group">
                                            <label>Password</label>
                                            <input type="password" placeholder="Password" class="form-control input-no-border"></input>
                                        </div>
                                    </div>
                                    <div class="footer text-center">
                                        <button type="submit" class="btn btn-fill btn-wd ">Let's go</button>
                                        <div class="forgot">
                                            <a href="#pablo">Forgot your password?</a>
                                        </div>
                                    </div>
                                </div>
                            </form>
                        </div>
                    </div>
                </div>
            </div>

        	<footer class="footer footer-transparent">
                <div class="container">
                    <div class="copyright">
                        &copy; 2017, made with <i class="fa fa-heart heart"></i> by <a href="http://www.gopheracademy.com">Gopher Academy</a>
                    </div>
                </div>
            </footer>
        </div>
    </div>
    </body>
        )
    }
}