/**
 * Retrieve information from the JARVIS endpoints periodically
 * @param {string} elementID The ID of the element to fill with data
 * @param {string} endpointURI The full URI of the API endpoint to poll
 * @param {int} everySeconds How often should the API endpoint be polled
 */
function getInfo(elementID, endpointURI, everySeconds) {
    
    // Initial Populate
    _getInfo(elementID, endpointURI);

    if (everySeconds > 0) {
        setInterval(function () {
            _getInfo(elementID, endpointURI);
        }, (everySeconds * 1000));
    } else {
        _getInfo(elementID, endpointURI);
    }
}

function refreshImage(elementID, everySeconds) {

    // Save base source    
    document.getElementById(elementID).setAttribute("data-source", document.getElementById(elementID).src);
    setInterval(function () { 
        d = new Date();
        var newSource = document.getElementById(elementID).getAttribute("data-source") + "?" + d.getTime()        
        document.getElementById(elementID).src = newSource;
    }, (everySeconds * 1000));
}

/**
 * Retrieve information from JARVIS endpoints 
 * @param {string} elementID The ID of the element to fill with data
 * @param {string} endpointURI The full URI of the API endpoint to poll
 */
function _getInfo(elementID, endpointURI) {
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.onreadystatechange = function() {
        if (xmlhttp.readyState == XMLHttpRequest.DONE ) {
            if (xmlhttp.status == 200) {
               document.getElementById(elementID).innerHTML = xmlhttp.responseText;
           }
        }
    }
    xmlhttp.open("GET", endpointURI, true);
    xmlhttp.send();
}

function HitAPI(endpointURI) {
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.open("GET", endpointURI, true);
    xmlhttp.send();
}


function getList(elementID, endpointURI, renderElement, everySeconds) {
    
    // Initial Populate
    _getList(elementID, endpointURI, renderElement);

    if (everySeconds > 0) {
        setInterval(function () {
            _getList(elementID, endpointURI, renderElement);
        }, (everySeconds * 1000));
    } else {
        _getList(elementID, endpointURI,renderElement);
    }
}
function _getList(elementID, endpointURI, renderElement) {
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.onreadystatechange = function() {
        if (xmlhttp.readyState == XMLHttpRequest.DONE ) {
            if (xmlhttp.status == 200) {
                var items = xmlhttp.responseText.split(",");
                var output = ""
                for (var i = 0; i < items.length; i++) {
                    output = output.concat(renderElement(items[i]))
                }
               document.getElementById(elementID).innerHTML = output
           }
        }
    }
    xmlhttp.open("GET", endpointURI, true);
    xmlhttp.send();
}