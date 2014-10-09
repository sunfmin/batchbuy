var webpack = require("webpack");

module.exports = {
	cache: true,
	entry: './react_src/Main.js',
	devtool: 'inline-source-map',
	output: {
		filename: './assets/bundle.js'
	},
	module: {
		loaders: [
			{test: /\.js$/, loader: 'jsx-loader?brfs'}
		]
	},
	plugins: [
		new webpack.IgnorePlugin(/vertx/)
	]
};
