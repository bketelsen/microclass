var m = require("mithril")

module.exports = {
    view: function() {
        return (
            <body>
            	<div class="wrapper">
		<div class="sidebar" data-background-color="brown" data-active-color="danger">
			<div class="logo">
				<a href="http://www.creative-tim.com" class="simple-text">
					Creative Tim
				</a>
			</div>
			<div class="logo logo-mini">
				<a href="http://www.creative-tim.com" class="simple-text">
					CT
				</a>
			</div>
	    	<div class="sidebar-wrapper">
				<div class="user">
	                <div class="photo">
	                    <img src="/assets/img/faces/face-2.jpg" />
	                </div>
	                <div class="info">
	                    <a data-toggle="collapse" href="#collapseExample" class="collapsed">
	                        Chet Faker
	                        <b class="caret"></b>
	                    </a>
	                    <div class="collapse" id="collapseExample">
	                        <ul class="nav">
	                            <li><a href="#profile">My Profile</a></li>
	                            <li><a href="#edit">Edit Profile</a></li>
	                            <li><a href="#settings">Settings</a></li>
	                        </ul>
	                    </div>
	                </div>
	            </div>
	            <ul class="nav">
					<li>
	                    <a data-toggle="collapse" href="#dashboardOverview">
	                        <i class="ti-panel"></i>
	                        <p>Dashboard
                                <b class="caret"></b>
                            </p>
	                    </a>
                        <div class="collapse" id="dashboardOverview">
							<ul class="nav">
								<li><a href="../dashboard/overview.html">Overview</a></li>
								<li><a href="../dashboard/stats.html">Stats</a></li>
							</ul>
						</div>
	                </li>
	                <li>
	                    <a data-toggle="collapse" href="#componentsExamples">
	                        <i class="ti-package"></i>
	                        <p>Components
	                           <b class="caret"></b>
	                        </p>
	                    </a>
	                    <div class="collapse" id="componentsExamples">
	                        <ul class="nav">
	                            <li><a href="../components/buttons.html">Buttons</a></li>
	                            <li><a href="../components/grid.html">Grid System</a></li>
	                            <li><a href="../components/panels.html">Panels</a></li>
	                            <li><a href="../components/sweet-alert.html">Sweet Alert</a></li>
	                            <li><a href="../components/notifications.html">Notifications</a></li>
	                            <li><a href="../components/icons.html">Icons</a></li>
	                            <li><a href="../components/typography.html">Typography</a></li>
	                        </ul>
	                    </div>
	                </li>
	                <li>
						<a data-toggle="collapse" href="#formsExamples">
	                        <i class="ti-clipboard"></i>
	                        <p>
								Forms
	                           <b class="caret"></b>
	                        </p>
	                    </a>
	                    <div class="collapse" id="formsExamples">
	                        <ul class="nav">
								<li><a href="../forms/regular.html">Regular Forms</a></li>
								<li><a href="../forms/extended.html">Extended Forms</a></li>
								<li><a href="../forms/validation.html">Validation Forms</a></li>
	                            <li><a href="../forms/wizard.html">Wizard</a></li>
	                        </ul>
	                    </div>
	                </li>
					<li>
						<a data-toggle="collapse" href="#tablesExamples">
	                        <i class="ti-view-list-alt"></i>
	                        <p>
								Table list
	                           <b class="caret"></b>
	                        </p>
	                    </a>
	                    <div class="collapse" id="tablesExamples">
	                        <ul class="nav">
								<li><a href="../tables/regular.html">Regular Tables</a></li>
	                            <li><a href="../tables/extended.html">Extended Tables</a></li>
	                            <li><a href="../tables/bootstrap-table.html">Bootstrap Table</a></li>
								<li><a href="../tables/datatables.net.html">DataTables.net</a></li>
	                        </ul>
	                    </div>
	                </li>
	                <li>
						<a data-toggle="collapse" href="#mapsExamples">
	                        <i class="ti-map"></i>
	                        <p>
								Maps
	                           <b class="caret"></b>
	                        </p>
	                    </a>
	                    <div class="collapse" id="mapsExamples">
	                        <ul class="nav">
								<li><a href="../maps/google.html">Google Maps</a></li>
								<li><a href="../maps/vector.html">Vector maps</a></li>
								<li><a href="../maps/fullscreen.html">Full Screen Map</a></li>
	                        </ul>
	                    </div>
	                </li>
	                <li>
	                    <a href="../calendar.html">
	                        <i class="ti-calendar"></i>
	                        <p>Calendar</p>
	                    </a>
	                </li>
	                <li>
	                    <a href="../charts.html">
	                        <i class="ti-bar-chart-alt"></i>
	                        <p>Charts</p>
	                    </a>
	                </li>
	                <li class="active">
						<a data-toggle="collapse" href="#pagesExamples" aria-expanded="true">
	                        <i class="ti-gift"></i>
	                        <p>
								Pages
	                           <b class="caret"></b>
	                        </p>
	                    </a>
	                    <div class="collapse in" id="pagesExamples">
	                        <ul class="nav">
								<li><a href="timeline.html">Timeline Page</a></li>
								<li class="active"><a href="user.html">User Page</a></li>
								<li><a href="login.html">Login Page</a></li>
								<li><a href="register.html">Register Page</a></li>
								<li><a href="lock.html">Lock Screen Page</a></li>
	                        </ul>
	                    </div>
	                </li>
	            </ul>
	    	</div>
	    </div>

	    <div class="main-panel">
			<nav class="navbar navbar-default">
	            <div class="container-fluid">
					<div class="navbar-minimize">
						<button id="minimizeSidebar" class="btn btn-fill btn-icon"><i class="ti-more-alt"></i></button>
					</div>
	                <div class="navbar-header">
	                    <button type="button" class="navbar-toggle">
	                        <span class="sr-only">Toggle navigation</span>
	                        <span class="icon-bar bar1"></span>
	                        <span class="icon-bar bar2"></span>
	                        <span class="icon-bar bar3"></span>
	                    </button>
	                    <a class="navbar-brand" href="#userpage">User Page</a>
	                </div>
	                <div class="collapse navbar-collapse">
						<form class="navbar-form navbar-left navbar-search-form" role="search">
	    					<div class="input-group">
	    						<span class="input-group-addon"><i class="fa fa-search"></i></span>
	    						<input type="text" value="" class="form-control" placeholder="Search..."></input>
	    					</div>
	    				</form>
	                    <ul class="nav navbar-nav navbar-right">
	                        <li>
	                            <a href="#stats" class="dropdown-toggle btn-magnify" data-toggle="dropdown">
	                                <i class="ti-panel"></i>
									<p>Stats</p>
	                            </a>
	                        </li>
	                        <li class="dropdown">
	                            <a href="#notifications" class="dropdown-toggle btn-rotate" data-toggle="dropdown">
	                                <i class="ti-bell"></i>
	                                <span class="notification">5</span>
									<p class="hidden-md hidden-lg">
										Notifications
										<b class="caret"></b>
									</p>
	                            </a>
	                            <ul class="dropdown-menu">
	                                <li><a href="#not1">Notification 1</a></li>
	                                <li><a href="#not2">Notification 2</a></li>
	                                <li><a href="#not3">Notification 3</a></li>
	                                <li><a href="#not4">Notification 4</a></li>
	                                <li><a href="#another">Another notification</a></li>
	                            </ul>
	                        </li>
							<li>
	                            <a href="#settings" class="btn-rotate">
									<i class="ti-settings"></i>
									<p class="hidden-md hidden-lg">
										Settings
									</p>
	                            </a>
	                        </li>
	                    </ul>
	                </div>
	            </div>
	        </nav>

	        <div class="content">
	            <div class="container-fluid">
	                <div class="row">
	                    <div class="col-lg-4 col-md-5">
	                        <div class="card card-user">
	                            <div class="image">
	                                <img src="../../assets/img/background.jpg" alt="..."/>
	                            </div>
	                            <div class="content">
	                                <div class="author">
	                                  <img class="avatar border-white" src="../../assets/img/faces/face-2.jpg" alt="..."/>
	                                  <h4 class="title">Chet Faker<br />
	                                     <a href="#"><small>@chetfaker</small></a>
	                                  </h4>
	                                </div>
	                                <p class="description text-center">
	                                    "I like the way you work it <br/>
	                                    No diggity <bl />
	                                    I wanna bag it up"
	                                </p>
	                            </div>
	                            <hr/>
	                            <div class="text-center">
	                                <div class="row">
	                                    <div class="col-md-3 col-md-offset-1">
	                                        <h5>12<br /><small>Files</small></h5>
	                                    </div>
	                                    <div class="col-md-4">
	                                        <h5>2GB<br /><small>Used</small></h5>
	                                    </div>
	                                    <div class="col-md-3">
	                                        <h5>24,6$<br /><small>Spent</small></h5>
	                                    </div>
	                                </div>
	                            </div>
	                        </div>
	                        <div class="card">
	                            <div class="header">
	                                <h4 class="title">Team Members</h4>
	                            </div>
	                            <div class="content">
	                                <ul class="list-unstyled team-members">
                                        <li>
                                            <div class="row">
                                                <div class="col-xs-3">
                                                    <div class="avatar">
                                                        <img src="../../assets/img/faces/face-0.jpg" alt="Circle Image" class="img-circle img-no-padding img-responsive"></img>
                                                    </div>
                                                </div>
                                                <div class="col-xs-6">
                                                    DJ Khaled
                                                    <br />
                                                    <span class="text-muted"><small>Offline</small></span>
                                                </div>
                                                <div class="col-xs-3 text-right">
                                                    <btn class="btn btn-sm btn-success btn-icon"><i class="fa fa-envelope"></i></btn>
                                                </div>
                                            </div>
                                        </li>
                                        <li>
                                            <div class="row">
                                                <div class="col-xs-3">
                                                    <div class="avatar">
                                                        <img src="../../assets/img/faces/face-1.jpg" alt="Circle Image" class="img-circle img-no-padding img-responsive"></img>
                                                    </div>
                                                </div>
                                                <div class="col-xs-6">
                                                    Creative Tim
                                                    <br />
                                                    <span class="text-success"><small>Available</small></span>
                                                </div>
                                                <div class="col-xs-3 text-right">
                                                    <btn class="btn btn-sm btn-success btn-icon"><i class="fa fa-envelope"></i></btn>
                                                </div>
                                            </div>
                                        </li>
                                        <li>
                                            <div class="row">
                                                <div class="col-xs-3">
                                                    <div class="avatar">
                                                        <img src="../../assets/img/faces/face-3.jpg" alt="Circle Image" class="img-circle img-no-padding img-responsive"></img>
                                                    </div>
                                                </div>
                                                <div class="col-xs-6">
                                                    Flume
                                                    <br />
                                                    <span class="text-danger"><small>Busy</small></span>
                                                </div>
                                                <div class="col-xs-3 text-right">
                                                    <btn class="btn btn-sm btn-success btn-icon"><i class="fa fa-envelope"></i></btn>
                                                </div>
                                            </div>
                                        </li>
	                                </ul>
	                            </div>
	                        </div>
	                    </div>
	                    <div class="col-lg-8 col-md-7">
	                        <div class="card">
	                            <div class="header">
	                                <h4 class="title">Edit Profile</h4>
	                            </div>
	                            <div class="content">
	                                <form>
	                                    <div class="row">
	                                        <div class="col-md-5">
	                                            <div class="form-group">
	                                                <label>Company</label>
	                                                <input type="text" class="form-control border-input" disabled placeholder="Company" value="Creative Code Inc."></input>
	                                            </div>
	                                        </div>
	                                        <div class="col-md-3">
	                                            <div class="form-group">
	                                                <label>Username</label>
	                                                <input type="text" class="form-control border-input" placeholder="Username" value="michael23"></input>
	                                            </div>
	                                        </div>
	                                        <div class="col-md-4">
	                                            <div class="form-group">
	                                                <label for="exampleInputEmail1">Email address</label>
	                                                <input type="email" class="form-control border-input" placeholder="Email"></input>
	                                            </div>
	                                        </div>
	                                    </div>
	                                    <div class="row">
	                                        <div class="col-md-6">
	                                            <div class="form-group">
	                                                <label>First Name</label>
	                                                <input type="text" class="form-control border-input" placeholder="Company" value="Chet"></input>
	                                            </div>
	                                        </div>
	                                        <div class="col-md-6">
	                                            <div class="form-group">
	                                                <label>Last Name</label>
	                                                <input type="text" class="form-control border-input" placeholder="Last Name" value="Faker"></input>
	                                            </div>
	                                        </div>
	                                    </div>
	                                    <div class="row">
	                                        <div class="col-md-12">
	                                            <div class="form-group">
	                                                <label>Address</label>
	                                                <input type="text" class="form-control border-input" placeholder="Home Address" value="Melbourne, Australia"></input>
	                                            </div>
	                                        </div>
	                                    </div>
	                                    <div class="row">
	                                        <div class="col-md-4">
	                                            <div class="form-group">
	                                                <label>City</label>
	                                                <input type="text" class="form-control border-input" placeholder="City" value="Melbourne"></input>
	                                            </div>
	                                        </div>
	                                        <div class="col-md-4">
	                                            <div class="form-group">
	                                                <label>Country</label>
	                                                <input type="text" class="form-control border-input" placeholder="Country" value="Australia"></input>
	                                            </div>
	                                        </div>
	                                        <div class="col-md-4">
	                                            <div class="form-group">
	                                                <label>Postal Code</label>
	                                                <input type="number" class="form-control border-input" placeholder="ZIP Code"></input>
	                                            </div>
	                                        </div>
	                                    </div>
	                                    <div class="row">
	                                        <div class="col-md-12">
	                                            <div class="form-group">
	                                                <label>About Me</label>
	                                                <textarea rows="5" class="form-control border-input" placeholder="Here can be your description" value="Mike">
Oh so, your weak rhyme
You doubt I'll bother, reading into it
I'll probably won't, left to my own devices
But that's the difference in our opinions.
													</textarea>
	                                            </div>
	                                        </div>
	                                    </div>
	                                    <div class="text-center">
	                                        <button type="submit" class="btn btn-info btn-fill btn-wd">Update Profile</button>
	                                    </div>
	                                    <div class="clearfix"></div>
	                                </form>
	                            </div>
	                        </div>
	                    </div>
	                </div>
	            </div>
	        </div>

	        <footer class="footer">
	            <div class="container-fluid">
	                <nav class="pull-left">
	                    <ul>
	                        <li>
	                            <a href="http://www.creative-tim.com">
	                                Creative Tim
	                            </a>
	                        </li>
	                        <li>
	                            <a href="http://blog.creative-tim.com">
	                               Blog
	                            </a>
	                        </li>
	                        <li>
	                            <a href="http://www.creative-tim.com/license">
	                                Licenses
	                            </a>
	                        </li>
	                    </ul>
	                </nav>
					<div class="copyright pull-right">
	                    &copy; 2017, made with <i class="fa fa-heart heart"></i> by <a href="http://www.creative-tim.com">Creative Tim</a>
	                </div>
	            </div>
	        </footer>
	    </div>
	</div>
    </body>
        ) 
    }
}