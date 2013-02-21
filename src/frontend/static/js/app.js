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

    /**
     * Initialize appliction
     */
    App.prototype.initialize = function() {
        this._$body
            .on('submit', 'form#frmVideoUrl', _.bind(this._handleVideoUrlSubmit, this))
            .on('submit', 'form#frmVideoEdito', _.bind(this._handleVideoEditSubmit, this))
            .on('change', '#rngStart, #rngEnd', _.debounce(_.bind(this._handleVideoRangeChange, this), 500));
    };

    /**
     * User wants to edit a video
     */
    App.prototype._handleVideoUrlSubmit = function(event) {
        event.preventDefault();

        var $element = $('input#videoUrl'),
            videoUrl = $element.val(),
            videoId;

        if ((videoId = this.parseVideoId(videoUrl)) === false) {
            alert('Could not parse video id :(');
            return;
        }

        this._loadYtVideo(videoId);
    };

    App.prototype._handleVideoEditSubmit = function(event) {

    };

    App.prototype._handleVideoRangeChange = function(event) {
        var $element = event.currentTarget,
            start = $('#rngStart').val(),
            end = $('#rngEnd').val();

        console.log($element);
        this._video.setPosition(start, end);
    };

    /**
     * Try to parse video id from input
     *
     * @param {String} videoUrl
     * @return {Boolean|String}
     */
    App.prototype.parseVideoId = function(videoUrl) {
        var videoMatch  = videoUrl.match(/v=([a-zA-Z0-9]{11})/);
        
        return videoMatch ? videoMatch[1] : false;
    };

    /**
     * Try to embed YT video into dom
     *
     * @param {String} videoId
     * @return {$.Deferred}
     */
    App.prototype._loadYtVideo = function(videoId) {
        this._video = new YTGifCreator.Video(videoId);

        $('#videoEdit').html(this._template.render('tpl-video-edit'));

        return this._video.load('ytplayer').then(_.bind(function() {
            $('#videoEditControls').html(
                this._template.render('tpl-video-edit-controls', {
                    duration: this._video.getDuration()
                })
            );
            
            this._video.play(true);
        }, this), function() {
            alert('Could not load YouTube video :(');
        });
    };
}(YTGifCreator));
