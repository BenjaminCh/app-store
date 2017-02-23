// Load inferno module
import {render} from 'inferno';

// Load routing modules
import { Router, Route } from 'inferno-router';
import createBrowserHistory from 'history/createBrowserHistory';

// Load app's components
import AppStoreApp from './AppStoreApp';
import MainSearchComponent from './MainSearchComponent';

// Load app's styles
import './styles/base.scss';

// Meant to place the index.html file to the dist folder
require('file-loader?name=[name].[ext]!../public/index.html');
// Same for materialize css font
require('file-loader?name=[name].[ext]!../src/fonts/material-icons-regular.woff2');

if (module.hot) {
    require('inferno-devtools');
}

const browserHistory = createBrowserHistory();

const routes = (
	<Router history={ browserHistory }>
		<Route component={ AppStoreApp }>
			<Route path="/" component={ MainSearchComponent } />
		</Route>
	</Router>
);

render(routes, document.getElementById('app'));

if (module.hot) {
    module.hot.accept()
}