(function(ns) {
    'use strict';

    var App = ns.App = function() {
        this._$body = $('body');

        this._api = new ns.Api();
    };

    App.prototype.initialize = function() {
        this._$body.on('submit', 'form#frmVideoUrl', _.bind(this._handleVideoUrlSubmit, this));
    };

    App.prototype._handleVideoUrlSubmit = function(event) {
        event.preventDefault();

        var $element = $('input#videoUrl'),
            videoUrl = $element.val();

        

        console.log(videoUrl);
    };
}(YTGifCreator));
