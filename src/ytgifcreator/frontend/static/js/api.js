(function(ns) {
    'use strict';

    var Api = ns.Api = function() {
        this._baseUrl = '/_api/';
    };

    /**
     * Make an Api call
     *
     * @param {String} action Name of the api call
     * @param {Object} data
     * @return {$.Deferred}
     */
    Api.prototype._doRequest = function(action, data) {
        if (!_.isString(action)) {
            throw new Error('Parameter action must have a value');
        }

        var dfd = new $.Deferred();

        $.ajax(this._baseUrl + action, {
            type: 'POST',
            data: data
        }).done(function() {
            dfd.resolve();
        }).fail(function() {
            dfd.reject();
        });

        return dfd.promise();
    };

    Api.prototype.requestGif = function(data) {
        return this._doRequest('RequestGif', data);
    };

}(YTGifCreator));
