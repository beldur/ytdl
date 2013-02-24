// Ben Almen's serialize Object Method
$.fn.serializeObject = function(){
    var obj = {};

    $.each( this.serializeArray(), function(i,o){
        var n = o.name,
        v = o.value;

        obj[n] = obj[n] === undefined ? v
            : $.isArray( obj[n] ) ? obj[n].concat( v )
            : [ obj[n], v ];
    });

    return obj;
};

$(document).ready(function() {
    'use strict';

    var app = new YTGifCreator.App();
    app.initialize();
});
