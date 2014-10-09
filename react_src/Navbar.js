/** @jsx React.DOM */
var React = require('react');

module.exports = React.createClass({
	displayName: "Navbar",
	render: function () {
		return (
			<div>
				<div id="logo">下午茶时间</div>
				<ul className="nav nav-pills">
					<li>
						<a href="/order.html">我的订单</a>
					</li>
					<li>
						<a href="/order_list.html">全员订单列表</a>
					</li>
					<li>
						<a href="/product.html">产品维护</a>
					</li>
					<li>
						<a href="/user_list.html">用户维护</a>
					</li>
					<li>
						<a href="/profile.html">我的信息</a>
					</li>
				</ul>
			</div>
		);
	}
});
