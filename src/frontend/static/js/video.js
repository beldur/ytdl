(function(ns) {
    'use strict';

    /**
     * Represents YouTube video
     *
     * @constructor
     * @param {String} videoId
     */
    var Video = ns.Video = function(videoId) {
        if (!videoId || videoId.length != 11) {
            throw new Error('Invalid video id!');
        }

        this._videoId = videoId;
        this._ref = null;
        this._loopCheckerHandle = null;
    };

    /**
     * Load Video into given container
     *
     * @todo add timeout
     * @param {String} container
     * @return {$.Deferred}
     */
    Video.prototype.load = function(container) {
        var dfd = new $.Deferred();

        if (document.getElementById(container) === null) {
            throw new Error("Can't find given container " + container);
        }

        swfobject.embedSWF("http://www.youtube.com/v/" + this._videoId + "?enablejsapi=1&playerapiid=ytplayer&version=3",
            container, "425", "356", "8", null, null, { allowScriptAccess: 'always' }, { id: 'ytVideo' });

        window.onYouTubePlayerReady = _.bind(function(playerId) {
            this._ref = document.getElementById('ytVideo');
            this._duration = this._ref.getDuration();

            dfd.resolve();
        }, this);

        return dfd.promise();
    };

    /**
     * Start playing the video
     *
     * @param {Boolean} loop
     * @param {Number} start Start position for playing/looping
     * @param {Number} end End position for playing/looping
     */
    Video.prototype.play = function(loop, start, end) {
        this.setLoop(loop)
            .setPosition(start || 0, end || this._duration);

        this._ref.playVideo();
    };

    /**
     * Set Video loop behavior
     *
     * @param {Boolean} loop
     * @chainable
     */
    Video.prototype.setLoop = function(loop) {
        this._loop = loop;

        if (this._loop) {
            this._loopCheckerHandle = setInterval(_.bind(this._loopChecker, this), 1000);
        } else {
            clearInterval(this._loopChecker);
        }

        return this;
    };

    /**
     * Checks current position and sets new position if necessary
     */
    Video.prototype._loopChecker = function() {
        if (this._end && Math.ceil(this._ref.getCurrentTime()) >= this._end) {
            this.setPosition(this._start);
        }
    };

    /**
     * Set current video start and end position
     *
     * @param {Number}
     * @chainable
     */
    Video.prototype.setPosition = function(start, end) {
        start = parseInt(start, 10);
        end = parseInt(end, 10);

        if (end) {
            if (start >= end) {
                throw new Error('End must be greater than statt!');
            }
            this._end = end;
        }

        this._start = start;
        this._ref.seekTo(start, true);

        return this;
    };

    Video.prototype.getDuration = function() {
        return this._duration;
    };

    Video.prototype.getVideoId = function() {
        return this._videoId;
    };

}(YTGifCreator));
