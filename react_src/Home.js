/** @jsx React.DOM */
var React = require('react');
var service = require('./Service');

module.exports = React.createClass({displayName: "Home",
	getInitialState: function () {
		return {products: []};
	},

	componentWillMount: function () {
		var self = this;
		service.ProductListOfDate("2014-10-08", function(products, err){
			self.setState({"products": products});
		});
	},

	render: function () {
		var productList = this.state.products.map(function(product){
			return(
				<li>{product.Name}</li>
				)
		})

		return (
			<div>
				{productList}
			</div>
			);
	}
});
