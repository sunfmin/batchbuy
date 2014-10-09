/** @jsx React.DOM */
var React = require('react');

var Navbar = require("./Navbar");

module.exports = React.createClass({displayName: "App",
	render: function () {
		return (
			<div>
				<div className="app">
					<Navbar />
					<this.props.activeRouteHandler/>
				</div>
			</div>

			);
	}
});
