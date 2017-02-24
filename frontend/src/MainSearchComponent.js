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
// Declare the search state manager
var helper = algoliasearchHelper(client, indexName, {
	disjunctiveFacets: ['category'] // setup facets on app.category
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
 * cleanSearch : Called when the search reset close button is clicked.
 * It clean reset the search and clean query text, facets and everything.
 */
function cleanSearch(instance, event) {
	// Clear UX
	JQuery('#query_input').val('');
	JQuery('input[data-facet-type=category]').attr('checked', false);
	
	// Clear algolia object
	helper.clearCache()
	.setQuery('')
	.clearRefinements();
	
	// Launch a ne search
	performSearch('');
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
}

/**
 * validURL : Returns true if the string looks like a valid URL.
 * Returns false otherwise.
 * TODO : Put it in an external service
 */
function validURL(str) {
	// From http://stackoverflow.com/questions/1701898/how-to-detect-whether-a-string-is-in-url-format-using-javascript
	var regexp = /(ftp|http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?/
   return regexp.test(str);
}

/**
 * getImaginaaryImage : Returns an image URL based on imaginaary service
 * Notes -
 * This is generated with an image proxy I built
 * It generates an url crafted from given params passed in the URL
 * (base image, facebook profile images, color, text, size, etc)
 * and perform operations on it (resize, crop, add a layer, merge images, change color, apply mask, etc.)
 * and stores it in a cache.
 * TODO : Put it in an external service
 */
function getImaginaaryImage(imageUrl, w, h, defaultBackgroundColor) {
	// Fallback image
	// the following URL will generate a color filled image with the blank text 'No image :(' on the top
	// 
	var baseImage = '';
	var cdnImage = 'http://imaginaary.com/image/fetch/';
	
	w = w || 200;
	h = h || 200;

	// Set the base image, is it a fill color or a fetched image?
	// If th estring is an URL, then set it as fetch image
	if (imageUrl && validURL(imageUrl) === true) {
		baseImage = 'img_' + encodeURIComponent(encodeURIComponent(imageUrl)); // Tell the backend the base is an image.
	}
	else {
		// Generate a default image (grey with a question mark on it)
		// http://imaginaary.com/image/fetch/w_80,h_80/t_%3F,fs_50,f_Roboto/bgc_eceff1
		baseImage = 't_%3F,fs_50,f_Roboto/bgc_' + (defaultBackgroundColor || 'eceff1') // Generate the default unknown image using a color fill background.
	}

	// Set URL
	cdnImage += 'w_' + w + ",";
	cdnImage += 'h_' + h + "/";
	cdnImage += baseImage;
	
    return cdnImage;
}

/**
 * renderHits : Renders hits list.
 * Render result list component in the given element based on query results.
 */
function renderHits(element, results) {
	// Render results (hits)
	JQuery(element).html(
		InfernoServer.renderToString(
			<div>
				{results.hits && results.hits.length > 0 && 
					<ul>
						{(results.hits || []).map(function(hit) {
							return (
								<li>
									<div class="row">
										<div class="col s2">
											<a href={hit.link} target="_blank">
												<img class="app-thumbmail" src={hit.image} alt={hit.name + " image"} width="70" height="70"/>
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
				}
			</div>
		)
	);
	// Attach image src loading error event handler.
	// Meant to replace broken link with a default image.
	JQuery('img.app-thumbmail').on('error', function(e) {
		JQuery(this).unbind("error").attr("src", getImaginaaryImage(null, 80, 80, 'eceff1'));
	});
}

/**
 * renderPagination : Renders the pagination.
 * Render the pagination component in the given element based on query results.
 */
function renderPagination(element, results) {
	// Render Pagination
	var currentPage = results.page;
	var minPage = 0;
	var maxPage = results.nbPages-1;
	
	JQuery(element).html(
		InfernoServer.renderToString(
			<div>
				{results.hits && results.hits.length > 0 && 
					<div class="center-align">
						<a href="#!" data-page={0} class={"waves-effect waves-light btn black-text grey lighten-5 " + ((maxPage > 1 && currentPage > 0) ? "" : "hide") }>
							{ "<<|" }
						</a>
						<a href="#!" data-page={currentPage-1} class={"waves-effect waves-light btn " + ((maxPage > 1 && currentPage > 0) ? "" : "hide") }>
							<i class="material-icons left">chevron_left</i>
							Prev
						</a>
						<a href="#!" data-page={currentPage+1} class={"waves-effect waves-light btn " + ((maxPage > 1 && currentPage + 1 <= maxPage) ? "" : "hide") }>
							<i class="material-icons right">chevron_right</i>
							Next
						</a>
						<a href="#!" data-page={maxPage} class={"waves-effect waves-light btn black-text grey lighten-5 " + ((maxPage > 1 && currentPage + 1 <= maxPage) ? "" : "hide") }>
							{ ">>|" }
						</a>
						<p>
							Page {currentPage+1} / {results.nbPages}
						</p>
					</div>
				}
			</div>
		)
	);
	// Attach on click event to pagination buttons
	JQuery(element).on('click', 'a[data-page]', function(e) {
		e.preventDefault();
		var page = JQuery(this).data('page');
		helper.setPage(page)
		.search();
	});
}

/**
 * renderBottomInformation : Renders bottom information section.
 * Render the bottom information section component in the given element based on query results.
 * Used to display something if the search returned no items.
 */
function renderBottomInformation(element, results) {
	// Render bottom information
	JQuery(element).html(
		InfernoServer.renderToString(
			<div>
				{!results.hits || results.hits.length == 0 && 
					<div class="center-align">
						<h5> Oups, no results found :( </h5>
						<p> No results were found with your filters. Test other filters.</p>
						<img src="http://i.giphy.com/2VHKqlpI3rqRG.gif" width="480" height="284" alt="No results found"/>
					</div>
				}
			</div>
		)
	);
}

/**
 * renderCategoryFacets : Renders category facets for the search.
 * Render the category facets component in the given element based on query results.
 */
function renderCategoryFacets (element, results) {
	JQuery(element).html(
		InfernoServer.renderToString(
			<div>
				{results.hits && results.hits.length > 0 && 
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
				}
			</div>
		)
	);
	// Attach on click event to facets checkboxes
	JQuery('input[data-facet-type=category]').on('click', function(e) {
		e.preventDefault();
		helper.toggleFacetRefinement(JQuery(this).data('facet-type'), JQuery(this).data('facet'))
		.search();
	});
}

/**
 * renderStats : Renders stats for the current search.
 * Render the search stats component in the given element based on query results.
 */
function renderStats (element, results) {
	JQuery(element).html(
		InfernoServer.renderToString(
			<span>
				<strong class="primary"> { results.nbHits } results found </strong> in <strong class="primary">{ results.processingTimeMS } </strong> ms
			</span>
		)
	);
}

/**
 * renderResults : Used to launch an Algolia search
 * and call the render function with the search results.
 */
function renderResults (results) {

	// Render hits
	renderHits('#filtered-apps', results);

	// Render bottom information
	renderBottomInformation('#bottom-information', results);

	// Render pagination
	renderPagination('#pagination', results);

	// Render results facets
	renderCategoryFacets('#facets', results);

	// Render query stats
	renderStats('#info-hits-time', results);
}

/**
 * MainSearchComponent : Simply load the view with the base form template.
 */
function MainSearchComponent() {

	// Attach the result event called every time an algolia search returns something.
	// Call the render function to render results on the page.
	helper.on('result', function(content) {
		renderResults(content);
	});

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
									<i class="material-icons black-text" onClick={ linkEvent(this, cleanSearch) }>close</i>
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
				<div id="bottom-information" class="col s12">
				</div>
			</div>
			<div class="row">
				<div id="pagination" class="col s10">
				</div>
			</div>
		</div>
	);
};