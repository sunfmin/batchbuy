/** @jsx React.DOM */

var React = require('react');
var Router = require('react-router');
var Route = Router.Route;
var DefaultRoute = Router.DefaultRoute;
var Routes = Router.Routes;
var Link = Router.Link;

var Home = require('./Home');
var App = require('./App');


var routes = (
	<Routes>
		<Route handler={App}>
			<DefaultRoute handler={Home}/>
		</Route>
	</Routes>
	);

React.renderComponent(routes, document.getElementById('lowteaapp'));

