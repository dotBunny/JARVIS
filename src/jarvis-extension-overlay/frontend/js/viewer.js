var bIsPanelOpen = false;
var bIsShowingStatsPanel = false;
var bIsShowingJIRAPanel = false;
var bIsShowingSpotifyPanel = false

$("li#button-spotify").on("click", function() {
  if ( !bIsShowingSpotifyPanel ) 
  {
    // Handle Panel
    if ( !bIsPanelOpen) {
      $("div#jarvis").animate({right: "0px"});
      bIsPanelOpen = true;
    }

    $(".active").removeClass("active");
    $(this).addClass("active");
    bIsShowingSpotifyPanel = true;
    $("div#panel-spotify").animate({right: "0px"});
    bIsShowingJIRAPanel = false;
    $("div#panel-jira").animate({right: String(-((window.innerWidth/100)*30)) + "px"});
    bIsShowingStatsPanel = false;
    $("div#panel-stats").animate({right: String(-((window.innerWidth/100)*30)) + "px"});   
  } 
  else 
  {
    bIsShowingSpotifyPanel = false;
    closePanel();
  }
});


$("li#button-stats").on("click", function() {
  if ( !bIsShowingStatsPanel ) 
  {
    // Handle Panel
    if ( !bIsPanelOpen) {
      $("div#jarvis").animate({right: "0px"});
      bIsPanelOpen = true;
    }

    $(".active").removeClass("active");
    $(this).addClass("active");
    bIsShowingStatsPanel = true;
    $("div#panel-stats").animate({right: "0px"});
    bIsShowingSpotifyPanel = false;
    $("div#panel-spotify").animate({right: String(-((window.innerWidth/100)*30)) + "px"});
    bIsShowingJIRAPanel = false;
    $("div#panel-jira").animate({right: String(-((window.innerWidth/100)*30)) + "px"});
  } 
  else 
  {
    bIsShowingStatsPanel = false;
    closePanel();
  }
});

$("li#button-jira").on("click", function() {
  if ( !bIsShowingJIRAPanel ) 
  {

    // Handle Panel
    if ( !bIsPanelOpen) {
      $("div#jarvis").animate({right: "0px"});
      bIsPanelOpen = true;
    }
    
    $(".active").removeClass("active");
    $(this).addClass("active");
    bIsShowingJIRAPanel = true;
    $("div#panel-jira").animate({right: "0px"});
    bIsShowingSpotifyPanel = false;
    $("div#panel-spotify").animate({right: String(-((window.innerWidth/100)*30)) + "px"});
    bIsShowingStatsPanel = false;
    $("div#panel-stats").animate({right: String(-((window.innerWidth/100)*30)) + "px"});
    
  } 
  else 
  {
    bIsShowingJIRAPanel = false;
    closePanel();
  }
});

$("div#cancel").on("click", function() {
  if ( bIsPanelOpen ) {
    closePanel();
  }
});


function closePanel()
{
  var amount = String(-((window.innerWidth/100)*30)) + "px";

  
  $("div#panel-spotify").animate({right: amount});
  $("div#panel-jira").animate({right: amount});
  $("div#panel-stats").animate({right: amount});

  $(".active").removeClass("active");
  
  $("div#jarvis").animate({right: amount});

  bIsPanelOpen = false;
  bIsShowingJIRAPanel = false;
  bIsShowingSpotifyPanel = false;
  bIsShowingStatsPanel = false;
}


closePanel();