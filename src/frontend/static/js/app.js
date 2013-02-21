(function(ns) {
    'use strict';

    /**
     * Application logic
     *
     * @constructor
     */
    var App = ns.App = function() {
        this._$body = $('body');

        this._template = new ns.Template();
        this._template.initialize();
        this._api = new ns.Api();
    };

    App.prototype.initialize = function() {
        this._$body.on('submit', 'form#frmVideoUrl', _.bind(this._handleVideoUrlSubmit, this));
    };

    App.prototype._handleVideoUrlSubmit = function(event) {
        event.preventDefault();

        var $element = $('input#videoUrl'),
            videoUrl = $element.val(),
            videoId;

        if ((videoId = this.parseVideoId(videoUrl)) === false) {
            alert('Could not parse video id :(');
            return;
        }

        var dfd = this._loadYtVideo(videoId).then(_.bind(function(videoInformation) {
            this._videoInformation = videoInformation;

            $('#videoEditControls').html(
                this._template.render('tpl-video-edit-controls', videoInformation)
            );
            
            console.log(videoInformation);
        }, this), function() {
            alert('Could not load YouTube video :(');
        });

    };

    App.prototype.parseVideoId = function(videoUrl) {
        var videoMatch  = videoUrl.match(/v=([a-zA-Z0-9]{11})/);
        
        return videoMatch ? videoMatch[1] : false;
    };

    /**
     * Try to embed YT video into dom
     *
     * @todo Add timeout
     * @param {String} videoId
     * @return {$.Deferred}
     */
    App.prototype._loadYtVideo = function(videoId) {
        var dfd = new $.Deferred();

        $('#videoEdit').html(this._template.render('tpl-video-edit'));

        swfobject.embedSWF("http://www.youtube.com/v/" + videoId + "?enablejsapi=1&playerapiid=ytplayer&version=3",
            "ytplayer", "425", "356", "8", null, null, { allowScriptAccess: 'always' }, { id: 'ytVideo' });

        window.onYouTubePlayerReady = function(playerId) {
            var ref = document.getElementById('ytVideo');

            dfd.resolve({
                reference: ref,
                duration: ref.getDuration()
            });
        };

        return dfd.promise();
    };
}(YTGifCreator));
