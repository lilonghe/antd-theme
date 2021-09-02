const path = require('path');
const webpack = require('webpack');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const CssMinimizerPlugin = require('css-minimizer-webpack-plugin');

const config = require('./config');

const cssSRC = {
    loader: "css-loader",
    options: {
        modules: {
            localIdentName: "[name]_[local]--[hash:base64:5]",
        },
        esModule: false,
    },
}

const libCSSSRC = {
    loader: "css-loader",
    options: {
        esModule: false,
    },
}

const lessSRC = {
    loader: 'less-loader',
    options: {
        lessOptions: {
            javascriptEnabled: true
        }
    }
};

module.exports = {
    mode: 'production',
    cache: {
        type: 'filesystem',
    },
    entry: {
        red: path.resolve(__dirname, '../src/themes', 'red.less'),
        default: path.resolve(__dirname, '../src/themes', 'default.less'),
    },
    output: {
		path: path.resolve(__dirname, '../public/themes'),
	},
   
    plugins: [
        new MiniCssExtractPlugin({
            filename: '[name].css',
        }),
    ],
    optimization: {
        minimize: true,
        minimizer: [
          new CssMinimizerPlugin(),
        ],
    },
    module: {
        rules: [
           
            {
                test: /\.less$/,
                include: /src/,
                use: [
                    MiniCssExtractPlugin.loader,
                    libCSSSRC,
                    lessSRC,
                ]
            },
            // assets loader
            {
                test: config.webpack.assetsPattern || /\.(png|jpg|gif|svg)$/i,
                use:[
                    {
                        loader: config.webpack.assetsLoader || 'url-loader',
                        options: config.webpack.assetsLoaderOption || {
                            limit: 1024,
                            outputPath: 'assets',
                        }
                    }
                ]
            },
        ]
    }
}