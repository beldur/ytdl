(function(ns) {
    'use strict';

    var Template = ns.Template = function() {
        this._templates = [];
        _.templateSettings.variable = 'data';
    };

    Template.prototype.initialize = function() {
        $('script[type="template/underscore"]').each(_.bind(function(index, item) {
            this._templates[item.id] = _.template(item.innerHTML);
        }, this));
    };

    Template.prototype.render = function(tplName, tplData) {
        return this._templates[tplName](tplData);
    };
}(YTGifCreator));
