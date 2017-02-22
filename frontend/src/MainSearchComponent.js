import Inferno, { linkEvent } from 'inferno';
import Component from 'inferno-component';
import InfernoServer from 'inferno-server';

// Expose the MainSearchComponent to the rest of the app.
export default MainSearchComponent;

/**
 * VARS declarations
 */

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
var helper = algoliasearchHelper(client, indexName);
var searchParams = {
	'highlightPreTag' : '<strong class="highlighted">',
	'highlightPostTag' : '</strong>',
	'facets' : 'category'
};

/**
 * FUNCTIONS declarations
 */

// Show export function first
// Simply load the view with the base form template.
function MainSearchComponent() {
	return (
		<div>
			<div class="nav-wrapper">
				<div class="row">
					<div class="input-field col s12">
						<label class="active" for="query_input">Search</label>
						<input id="query_input" type="text" onKeyUp={ linkEvent(this, queryTextChanged) }/>
					</div>
				</div>
			</div>
			<div class="row">
				<div id="info-hits-time" class="col s2"> </div>
			</div>
			<div class="row">
				<div id="filtered-apps" class="col s9">
				</div>
				<div id="facets" class="col s3">
				</div>
			</div>
		</div>
	);
};

/**
 * queryTextChanged : Called when the query text change.
 * It calls the performSearch function to launch a search.
 */
function queryTextChanged(instance, event) {
	performSearch(event.srcElement.value);
}

function facetItemClicked(instance, event) {

}

function performSearch(query) {
	index.search(query, searchParams)
	.then(function (content) {
		renderResults(content);
	});
}

/**
 * renderResults : Used to launch an Algolia search
 * and call the render function with the search results.
 */
function renderResults (results) {
	// Render results
	document.getElementById('filtered-apps').innerHTML = InfernoServer.renderToString(
		<div>
			<div class="row">
				<div class="col s2 offset-s10"> </div>
			</div>
			<ul>
				{results.hits.map(function(hit) {
					return (
						<li dangerouslySetInnerHTML={{__html: hit._highlightResult.name.value}}></li>
					);
				})}
			</ul>
		</div>
	);

	// Render results facets
	document.getElementById('facets').innerHTML = InfernoServer.renderToString(
		<div>
			{Object.keys(results.facets['category']).map(function(facet) {
				return (
					<p>
						<input type="checkbox" id={"checkbox-" + facet} onClick={ linkEvent(this, facetItemClicked) }/>
						<label for={"checkbox-" + facet}>{facet} ({results.facets['category'][facet]}) </label>
					</p>
				);
			})}
		</div>
	);

	// Render query stats
	document.getElementById('info-hits-time').innerHTML = InfernoServer.renderToString(
		(<span>
			<strong class="primary"> { results.nbHits }</strong> hits in <strong class="primary">{ results.processingTimeMS } </strong> ms
		</span>)
	);
}