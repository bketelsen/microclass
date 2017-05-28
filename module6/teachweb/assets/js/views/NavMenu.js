var m = require("mithril")

module.exports = {
    view: function() {
        return  (
 
 <nav class="navbar navbar-transparent navbar-absolute">
        <div class="container">
            <div class="navbar-header">
                <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#navigation-example-2">
                    <span class="sr-only">Toggle navigation</span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                </button>
                <a class="navbar-brand" href="../dashboard/overview.html">Gopher Academy Online</a>
            </div>
            <div class="collapse navbar-collapse">
                <ul class="nav navbar-nav navbar-right">
                    <li>
                       <a href="register.html">
                            Register
                        </a>
                    </li>
					<li>
                       <a href="../dashboard/overview.html">
                            Dashboard
                        </a>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
            )
    }
}