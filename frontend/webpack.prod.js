const webpack = require('webpack');
const path = require('path');

const ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
	entry: './src/index.js',
	output: {
		path: './dist/',
		filename: 'bundle.js',
		publicPath: 'http://search.tests.sh/'
	},
	module: {
		loaders: [
			{
				test: /\.js?/,
				exclude: [/node_modules/, /styles/],
				loaders: ['babel-loader'],
				include: path.join(__dirname, 'src')
			},
			{
				test: /\.scss$/,
				loader: 'style-loader!css-loader!sass-loader!resolve-url-loader!sass-loader?sourceMap'
			},
			{
				test: /\.css$/,
				loader: "style-loader!css-loader"
			},
			{
				test: /\.woff(2)?(\?v=[0-9]\.[0-9]\.[0-9])?$/,
				loader: 'url-loader?limit=10000&mimetype=application/font-woff'
			},
			{
				test: /\.(ttf|eot|svg)(\?v=[0-9]\.[0-9]\.[0-9])?$/,
				loader: 'file-loader'
			}
        ]
	},
	devServer: {
		contentBase: './',
		port: 8080,
		noInfo: false,
		hot: true,
		inline: true,
		proxy: {
			'/': {
				bypass: function (req, res, proxyOptions) {
					return '/public/index.html';
				}
			}
		}
	},
	plugins: [
		new webpack.HotModuleReplacementPlugin(),
		new ExtractTextPlugin('public/main.css')
	],
	resolve: {
		extensions: ['.js', '.css', '.scss', 'woff', 'woff2', 'ttf', 'eot', 'svg'],
		alias: {
			materialize: path.join(__dirname, '/node_modules/materialize-css/sass/materialize.scss'),
		}
	}
};