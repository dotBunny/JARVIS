/*
Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License"). You may not use this file except in compliance with the License. A copy of the License is located at

    http://aws.amazon.com/apache2.0/

or in the "license" file accompanying this file. This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
*/

/*

  Set Javascript specific to the extension viewer view in this file.

*/

var bIsPanelOpen = false;



function panelClose()
{
  var amount = -((window.innerWidth/100)*30);
  $("div#jarvis-panel").animate({right:  String(amount) + "px"});
  bIsPanelOpen = false;
}

function panelOpen()
{
  $("div#jarvis-panel").animate({right: "0px"});
  bIsPanelOpen = open;

  // Add close catcher to window
}

function panelToggle()
{
  if ( bIsPanelOpen ) {
    panelClose();
  } else { 
    panelOpen();
  }
}

$( "div#jarvis-button" ).on( "click", function() {
  panelToggle();
});



