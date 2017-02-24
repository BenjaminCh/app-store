import Inferno, { linkEvent } from 'inferno';
import Component from 'inferno-component';
import InfernoServer from 'inferno-server';
import JQuery from 'jquery'

// Show exported component first
// Expose the MainSearchComponent to the rest of the app.
export default MainSearchComponent;

/**
 * VARS declarations
 */

// TODO : Move those into an External Service

// Load Algolia Search & Helper
var algoliasearch = require('algoliasearch');
var algoliasearchHelper = require('algoliasearch-helper')

// Initialize algolia client and index for search
// TODO : Declare keys elsewhere
var applicationID = 'SFKD8Z2O1Y';
var apiKey = '5714bc93d8c05722a95ca84081f01e96';
var indexName = 'apps';
var client = algoliasearch(applicationID, apiKey);
var index = client.initIndex(indexName);
var helper = algoliasearchHelper(client, indexName, {
	disjunctiveFacets: ['category']
});

/**
 * FUNCTIONS declarations
 */

/**
 * queryTextChanged : Called when the query text change.
 * It calls the performSearch function to launch a search.
 */
function queryTextChanged(instance, event) {
	performSearch(event.srcElement.value);
}

/**
 * performSearch : Perform a search and call the rendering once done.
 */
function performSearch(query) {
	helper
	.setQuery(query)
	.setQueryParameter('highlightPreTag', '<strong class="highlighted">')
	.setQueryParameter('highlightPostTag', '</strong>')
	.setQueryParameter('hitsPerPage', 10)
    .search();

	location.replace('#q=' + encodeURIComponent(query));
}

function validURL(str) {
	// From http://stackoverflow.com/questions/1701898/how-to-detect-whether-a-string-is-in-url-format-using-javascript
	var regexp = /(ftp|http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?/
   return regexp.test(str);
}

function getImageUrl(imageUrl, w, h, defaultBackgroundColor) {
	// Fallback image
	// the following URL will generate a color filled image with the blank text 'No image :(' on the top
	// Notes -
	// This is generated with an image proxy I built
	// It can perform operations on images and allow caching on it
	// It can build images from Facebook profile picture, fetched images from an url
	// and perform operations on it (resize, crop, add a layer, merge images, etc.)
	// TODO : Put it in an external service
	var baseImage = '';
	var cdnImage = 'http://imaginaary.com/image/fetch/';
	
	w = w || 200;
	h = h || 200;

	if (imageUrl && validURL(imageUrl) === true) {
		baseImage = 'img_' + encodeURIComponent(encodeURIComponent(imageUrl)); // Tell the backend the base is an image.
	}
	else {
		baseImage = 't_%3F,fs_50,f_Roboto/bgc_' + (defaultBackgroundColor || 'eceff1') // Generate the default unknown image using a color fill background.
	}

	cdnImage += 'w_' + w + ",";
	cdnImage += 'h_' + h + "/";
	cdnImage += baseImage;
	
    return cdnImage;
}

/**
 * renderResults : Used to launch an Algolia search
 * and call the render function with the search results.
 */
function renderResults (results) {

	// Render results
	document.getElementById('filtered-apps').innerHTML = InfernoServer.renderToString(
		<div>
			<ul>
				{(results.hits || []).map(function(hit) {
					return (
						<li>
							<div class="row">
								<div class="col s2">
									<a href={hit.link} target="_blank">
										<img class="app-thumbmail" src={ getImageUrl(hit.image, (hit.name + " image"), 70, 70)} alt={hit.name + " image"} width="70" height="70"/>
									</a>
								</div>
								<div class="col s10">
									<h6 dangerouslySetInnerHTML={{__html: hit._highlightResult.name.value}}></h6>
									<p>
										<i dangerouslySetInnerHTML={{__html: hit._highlightResult.category.value}}></i>
									</p>
								</div>
							</div>
						</li>
					);
				})}
			</ul>
		</div>
	);
	// Attach image src loading error event handler.
	// Meant to replace broken link with a default image.
	JQuery('img.app-thumbmail').on('error', function(e) {
		JQuery(this).unbind("error").attr("src", getImageUrl(null, 80, 80, 'eceff1'));
	});

	// Render Pagination
	var currentPage = results.page;
	var minPage = 0;
	var maxPage = results.nbPages-1;
	
	document.getElementById('pagination').innerHTML = InfernoServer.renderToString(
		<div class="center-align">
			<a href="#!" data-page={0} class={"waves-effect waves-light btn black-text grey lighten-5 " + ((maxPage > 1 && currentPage > 0) ? "" : "hide") }>
				{ "<<|" }
			</a>
			<a href="#!" data-page={currentPage-1} class={"waves-effect waves-light btn " + ((maxPage > 1 && currentPage > 0) ? "" : "hide") }>
				<i class="material-icons left">chevron_left</i>
				Prev
			</a>
			<a href="#!" data-page={currentPage+1} class={"waves-effect waves-light btn " + ((maxPage > 1 && currentPage + 2 <= maxPage) ? "" : "hide") }>
				<i class="material-icons right">chevron_right</i>
				Next
			</a>
			<a href="#!" data-page={maxPage} class={"waves-effect waves-light btn black-text grey lighten-5 " + ((maxPage > 1 && currentPage + 2 <= maxPage) ? "" : "hide") }>
				{ ">>|" }
			</a>
		</div>
	);
	// Attach on click event to pagination buttons
	JQuery('a[data-page]').on('click', function(e) {
		e.preventDefault();
		helper.setPage(JQuery(this).data('page'))
		.search();
	});

	// Render results facets
	document.getElementById('facets').innerHTML = InfernoServer.renderToString(
		<div>
			<h5> Categories </h5>
			{(results.getFacetValues('category', {sortBy: ['count:desc']}) || []).map(function(facet, index) {
				return (
					<p>
						<input type="checkbox" id={"checkbox-facet-" + index} data-facet={facet.name} data-facet-type="category" checked={facet.isRefined}/>
						<label for={"checkbox-facet-" + index}>{facet.name} ({facet.count}) </label>
					</p>
				);
			})}
		</div>
	);
	// Attach on click event to facets checkboxes
	JQuery('input[data-facet-type=category]').on('click', function(e) {
		e.preventDefault();
		helper.toggleFacetRefinement(JQuery(this).data('facet-type'), JQuery(this).data('facet'))
		.search();
	});

	// Render query stats
	document.getElementById('info-hits-time').innerHTML = InfernoServer.renderToString(
		(<span>
			<strong class="primary"> { results.nbHits } results found </strong> in <strong class="primary">{ results.processingTimeMS } </strong> ms
		</span>)
	);
}

// Simply load the view with the base form template.
function MainSearchComponent() {

	// Attach the result event called every time an algolia search returns something.
	// Call the render function to render results on the page.
	helper.on('result', function(content) {
		renderResults(content);
	});

	// Look for any initial query in the url
	// Inspired by https://github.com/algolia/algoliasearch-client-javascript/blob/master/examples/instantsearch%2Bpagination.html 
	if (location.hash && location.hash.indexOf('#q=') === 0) {
		var params = location.hash.substring(3);
		var pageParamOffset = params.indexOf('&page=');
		var q, page;
		if (pageParamOffset > -1) {
			q = decodeURIComponent(params.substring(0, pageParamOffset));
			JQuery('#query_input').val(q);
			page = params.substring(pageParamOffset).split('=')[1];
		}
		else {
			q = decodeURIComponent(params);
			page = 1;
		}
	}

	// Launch a search to initialize results list
	// performSearch(q);

	return (
		<div>
			<div class="row">
				<div class="input-field col s12">
					<nav class="grey lighten-5 black-text text-darken-2">
						<div class="nav-wrapper">
							<form>
								<div class="input-field">
									<input id="query_input" type="search" onKeyUp={ linkEvent(this, queryTextChanged) } placeholder="Search for apps"/>
									<label class="label-icon" for="search"><i class="material-icons black-text">search</i></label>
									<i class="material-icons black-text">close</i>
								</div>
							</form>
						</div>
					</nav>
				</div>
			</div>
			<div class="row">
				<div id="info-hits-time" class="col s4"> </div>
				<div id="typos" class="col s4"> </div>
				<div id="ranking" class="col s4"> </div>
			</div>
			<div class="row">
				<div id="filtered-apps" class="col s9">
				</div>
				<div id="facets" class="col s3">
				</div>
			</div>
			<div class="row">
				<div id="pagination" class="col s10">
				</div>
			</div>
		</div>
	);
};