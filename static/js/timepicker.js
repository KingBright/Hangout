(function (global, factory) {
    if (typeof define === "function" && define.amd) {
        define(['exports'], factory);
    } else if (typeof exports !== "undefined") {
        factory(exports);
    } else {
        var mod = {
            exports: {}
        };
        factory(mod.exports);
        global.mdTimePicker = mod.exports;
    }
})(this, function (exports) {
    'use strict';

    Object.defineProperty(exports, "__esModule", {
        value: true
    });

    function _classCallCheck(instance, Constructor) {
        if (!(instance instanceof Constructor)) {
            throw new TypeError("Cannot call a class as a function");
        }
    }

    var _createClass = function () {
            function defineProperties(target, props) {
                for (var i = 0; i < props.length; i++) {
                    var descriptor = props[i];
                    descriptor.enumerable = descriptor.enumerable || !1;
                    descriptor.configurable = !0;
                    if ("value" in descriptor) descriptor.writable = !0;
                    Object.defineProperty(target, descriptor.key, descriptor);
                }
            }

            return function (Constructor, protoProps, staticProps) {
                if (protoProps) defineProperties(Constructor.prototype, protoProps);
                if (staticProps) defineProperties(Constructor, staticProps);
                return Constructor;
            };
        }(),
        _dialog = {
            view: !0,
            state: !1
        },
        mdTimePicker = function () {
            /**
             * [constructor of the mdTimePicker]
             *
             * @method constructor
             *
             * @param  {moment}   future                                                        [the future moment till which the calendar shall render] [@default = init]
             * @param    {Boolean}  mode                                                                        [this value tells whether the time dialog will have the 24 hour mode (true) or 12 hour mode (false)] [@default = false]
             * @param  {String}   orientation = 'LANDSCAPE' or 'PORTRAIT'  [force the orientation of the picker @default = 'LANDSCAPE']
             * @param  {String}  ok = 'ok'                                                                    [ok button's text]
             * @param  {String}  cancel = 'cancel'                                                    [cancel button's text]
             * @param  {Boolean} colon = true                                                            [add an option to enable quote in 24 hour mode]
             *
             * @return {Object}                                                                                    [mdTimePicker]
             */

            function mdTimePicker(_ref) {
                _ref = {}
                var type = 'time',
                    _ref$init = _ref.init,
                    init = _ref$init === undefined ? moment() : _ref$init,
                    _ref$orientation = _ref.orientation,
                    orientation = _ref$orientation === undefined ? 'LANDSCAPE' : _ref$orientation,
                    _ref$ok = _ref.ok,
                    ok = _ref$ok === undefined ? 'ok' : _ref$ok,
                    _ref$cancel = _ref.cancel,
                    cancel = _ref$cancel === undefined ? 'cancel' : _ref$cancel,
                    _ref$colon = _ref.colon,
                    colon = _ref$colon === undefined ? !0 : _ref$colon;

                _classCallCheck(this, mdTimePicker);

                this._type = type;
                this._init = init;
                this._orientation = orientation;
                this._ok = ok;
                this._cancel = cancel;
                this._colon = colon;

                /**
                 * [dialog selected classes have the same structure as dialog but one level down]
                 * @type {Object}
                 * e.g
                 * sDialog = {
   *   picker: 'some-picker-selected'
   * }
                 */
                this._sDialog = {};
                // attach the dialog if not present
                if (!document.getElementById('mddtp-picker__' + this._type)) {
                    this._buildDialog();
                }
            }

            /**
             * [time to get or set the current picker's moment]
             *
             * @method time
             *
             * @param  {moment} m
             *
             */


            _createClass(mdTimePicker, [{
                key: 'toggle',
                value: function toggle() {
                    this._selectDialog();
                    // work according to the current state of the dialog
                    if (mdTimePicker.dialog.state) {
                        this._hideDialog();
                    } else {
                        this._initTimeDialog(this._init);
                        this._showDialog();
                    }
                }
            }, {
                key: '_selectDialog',
                value: function _selectDialog() {
                    // now do what you normally would do
                    this._sDialog.picker = document.getElementById('mddtp-picker__' + [this._type]);
                    /**
                     * [sDialogEls stores all inner components of the selected dialog or sDialog to be later getElementById]
                     *
                     * @type {Array}
                     */
                    var sDialogEls = ['header', 'cancel', 'ok', 'left', 'right', 'subtitle', 'title', 'AM', 'PM', 'needle', 'hourView', 'minuteView', 'hour', 'minute', 'fakeNeedle', 'circularHolder', 'circle', 'dotSpan'],
                        i = sDialogEls.length;

                    while (i--) {
                        this._sDialog[sDialogEls[i]] = document.getElementById('mddtp-' + this._type + '__' + sDialogEls[i]);
                    }

                    this._sDialog.tDate = this._init.clone();
                    this._sDialog.sDate = this._init.clone();
                }
            }, {
                key: '_showDialog',
                value: function _showDialog() {
                    var me = this,
                        zoomIn = 'zoomIn';

                    mdTimePicker.dialog.state = !0;
                    this._sDialog.picker.classList.remove('mddtp-picker--inactive');
                    this._sDialog.picker.classList.add(zoomIn);
                    // if the dialog is forced into portrait mode
                    if (this._orientation === 'PORTRAIT') {
                        this._sDialog.picker.classList.add('mddtp-picker--portrait');
                    }
                    setTimeout(function () {
                        me._sDialog.picker.classList.remove(zoomIn);
                    }, 300);
                }
            }, {
                key: '_hideDialog',
                value: function _hideDialog() {
                    var me = this,
                        subtitle = me._sDialog.subtitle,
                        AM = this._sDialog.AM,
                        PM = this._sDialog.PM,
                        minute = this._sDialog.minute,
                        hour = this._sDialog.hour,
                        minuteView = this._sDialog.minuteView,
                        hourView = this._sDialog.hourView,
                        picker = this._sDialog.picker,
                        needle = this._sDialog.needle,
                        dotSpan = this._sDialog.dotSpan,
                        active = 'mddtp-picker__color--active',
                        inactive = 'mddtp-picker--inactive',
                        zoomIn = 'zoomIn',
                        zoomOut = 'zoomOut',
                        hidden = 'mddtp-picker__circularView--hidden',
                        selection = 'mddtp-picker__selection';

                    mdTimePicker.dialog.state = !1;
                    mdTimePicker.dialog.view = !0;
                    this._sDialog.picker.classList.add(zoomOut);
                    // reset classes
                    AM.classList.remove(active);
                    PM.classList.remove(active);
                    minute.classList.remove(active);
                    hour.classList.add(active);
                    minuteView.classList.add(hidden);
                    hourView.classList.remove(hidden);
                    subtitle.setAttribute('style', 'display: none');
                    dotSpan.setAttribute('style', 'display: none');
                    needle.className = selection;

                    setTimeout(function () {
                        // remove portrait mode
                        me._sDialog.picker.classList.remove('mddtp-picker--portrait');
                        me._sDialog.picker.classList.remove(zoomOut);
                        me._sDialog.picker.classList.add(inactive);
                        // clone elements and add them again to clear events attached to them
                        var pickerClone = picker.cloneNode(!0);
                        picker.parentNode.replaceChild(pickerClone, picker);
                    }, 300);
                }
            }, {
                key: '_buildDialog',
                value: function _buildDialog() {
                    var type = this._type,
                        docfrag = document.createDocumentFragment(),
                        container = document.createElement('div'),
                        header = document.createElement('div'),
                        body = document.createElement('div'),
                        action = document.createElement('div'),
                        cancel = document.createElement('button'),
                        ok = document.createElement('button');
                    // outer most container of the picker

                    // header container of the picker

                    // body container of the picker

                    // action elements container

                    // ... add properties to them
                    container.id = 'mddtp-picker__' + type;
                    container.classList.add('mddtp-picker');
                    container.classList.add('mddtp-picker-' + type);
                    container.classList.add('mddtp-picker--inactive');
                    container.classList.add('animated');
                    this._addId(header, 'header');
                    this._addClass(header, 'header');
                    // add header to container
                    container.appendChild(header);
                    this._addClass(body, 'body');
                    body.appendChild(action);
                    // add body to container
                    container.appendChild(body);
                    // add stuff to header and body according to dialog type
                    var _title = document.createElement('div'),
                        hour = document.createElement('span'),
                        span = document.createElement('span'),
                        minute = document.createElement('span'),
                        _subtitle = document.createElement('div'),
                        AM = document.createElement('div'),
                        PM = document.createElement('div'),
                        circularHolder = document.createElement('div'),
                        needle = document.createElement('div'),
                        dot = document.createElement('span'),
                        line = document.createElement('span'),
                        circle = document.createElement('span'),
                        minuteView = document.createElement('div'),
                        fakeNeedle = document.createElement('div'),
                        hourView = document.createElement('div');

                    // add properties to them
                    // inside header
                    this._addId(_title, 'title');
                    this._addClass(_title, 'title');
                    this._addId(hour, 'hour');
                    hour.classList.add('mddtp-picker__color--active');
                    span.textContent = ':';
                    this._addId(span, 'dotSpan');
                    span.setAttribute('style', 'display: none');
                    this._addId(minute, 'minute');
                    this._addId(_subtitle, 'subtitle');
                    this._addClass(_subtitle, 'subtitle');
                    _subtitle.setAttribute('style', 'display: none');
                    this._addId(AM, 'AM');
                    AM.textContent = 'AM';
                    this._addId(PM, 'PM');
                    PM.textContent = 'PM';
                    // add them to title and subtitle
                    _title.appendChild(hour);
                    _title.appendChild(span);
                    _title.appendChild(minute);
                    _subtitle.appendChild(AM);
                    _subtitle.appendChild(PM);
                    // add them to header
                    header.appendChild(_title);
                    header.appendChild(_subtitle);
                    // inside body
                    this._addId(circularHolder, 'circularHolder');
                    this._addClass(circularHolder, 'circularHolder');
                    this._addId(needle, 'needle');
                    needle.classList.add('mddtp-picker__selection');
                    this._addClass(dot, 'dot');
                    this._addClass(line, 'line');
                    this._addId(circle, 'circle');
                    this._addClass(circle, 'circle');
                    this._addId(minuteView, 'minuteView');
                    minuteView.classList.add('mddtp-picker__circularView');
                    minuteView.classList.add('mddtp-picker__circularView--hidden');
                    this._addId(fakeNeedle, 'fakeNeedle');
                    fakeNeedle.classList.add('mddtp-picker__circle--fake');
                    this._addId(hourView, 'hourView');
                    hourView.classList.add('mddtp-picker__circularView');
                    // add them to needle
                    needle.appendChild(dot);
                    needle.appendChild(line);
                    needle.appendChild(circle);
                    // add them to circularHolder
                    circularHolder.appendChild(needle);
                    circularHolder.appendChild(minuteView);
                    circularHolder.appendChild(fakeNeedle);
                    circularHolder.appendChild(hourView);
                    // add them to body
                    body.appendChild(circularHolder);

                    action.classList.add('mddtp-picker__action');
                    this._addId(cancel, 'cancel');
                    cancel.classList.add('mddtp-button');
                    cancel.setAttribute('type', 'button');
                    this._addId(ok, 'ok');
                    ok.classList.add('mddtp-button');
                    ok.setAttribute('type', 'button');
                    // add actions
                    action.appendChild(cancel);
                    action.appendChild(ok);
                    // add actions to body
                    body.appendChild(action);
                    docfrag.appendChild(container);
                    // add the container to the end of body
                    document.getElementsByTagName('body').item(0).appendChild(docfrag);
                }
            }, {
                key: '_initTimeDialog',
                value: function _initTimeDialog(m) {
                    var hour = this._sDialog.hour,
                        minute = this._sDialog.minute,
                        dotSpan = this._sDialog.dotSpan;

                    var text = parseInt(m.format('H'), 10);
                    if (text === 0) {
                        text = '00';
                    }
                    this._fillText(hour, text);
                    // add the configurable colon in this mode issue #56
                    if (this._colon) {
                        dotSpan.removeAttribute('style');
                    }

                    this._fillText(minute, m.format('mm'));
                    this._initHour();
                    this._initMinute();
                    this._attachEventHandlers();
                    this._dragDial();
                    this._switchToView(hour);
                    this._switchToView(minute);
                    this._addClockEvent();
                    this._setButtonText();
                }
            }, {
                key: '_initHour',
                value: function _initHour() {
                    var hourView = this._sDialog.hourView,
                        needle = this._sDialog.needle,
                        hour = 'mddtp-hour__selected',
                        selected = 'mddtp-picker__cell--selected',
                        rotate = 'mddtp-picker__cell--rotate-',
                        cell = 'mddtp-picker__cell',
                        docfrag = document.createDocumentFragment(),

                        hourNow = parseInt(this._sDialog.tDate.format('H'), 10);
                    for (var i = 1, j = 5; i <= 24; i++, j += 5) {
                        var div = document.createElement('div'),
                            span = document.createElement('span');

                        div.classList.add(cell);
                        // CHANGED exception case for 24 => 0 issue #57
                        if (i === 24) {
                            span.textContent = '00';
                        } else {
                            span.textContent = i;
                        }
                        div.classList.add(rotate + j);
                        if (hourNow === i) {
                            div.id = hour;
                            div.classList.add(selected);
                            needle.classList.add(rotate + j);
                        }
                        // CHANGED exception case for 24 => 0 issue #58
                        if (i === 24 && hourNow === 0) {
                            div.id = hour;
                            div.classList.add(selected);
                            needle.classList.add(rotate + j);
                        }
                        div.appendChild(span);
                        docfrag.appendChild(div);
                    }

                    //empty the hours
                    while (hourView.lastChild) {
                        hourView.removeChild(hourView.lastChild);
                    }
                    // set inner html accordingly
                    hourView.appendChild(docfrag);
                }
            }, {
                key: '_initMinute',
                value: function _initMinute() {
                    var minuteView = this._sDialog.minuteView,
                        minuteNow = parseInt(this._sDialog.tDate.format('m'), 10),
                        sMinute = 'mddtp-minute__selected',
                        selected = 'mddtp-picker__cell--selected',
                        rotate = 'mddtp-picker__cell--rotate-',
                        cell = 'mddtp-picker__cell',
                        docfrag = document.createDocumentFragment();

                    for (var i = 5, j = 10; i <= 60; i += 5, j += 10) {
                        var div = document.createElement('div'),
                            span = document.createElement('span');

                        div.classList.add(cell);
                        if (i === 60) {
                            span.textContent = this._numWithZero(0);
                        } else {
                            span.textContent = this._numWithZero(i);
                        }
                        if (minuteNow === 0) {
                            minuteNow = 60;
                        }
                        div.classList.add(rotate + j);
                        // (minuteNow === 1 && i === 60) for corner case highlight 00 at 01
                        if (minuteNow === i || minuteNow - 1 === i || minuteNow + 1 === i || minuteNow === 1 && i === 60) {
                            div.id = sMinute;
                            div.classList.add(selected);
                        }
                        div.appendChild(span);
                        docfrag.appendChild(div);
                    }
                    //empty the hours
                    while (minuteView.lastChild) {
                        minuteView.removeChild(minuteView.lastChild);
                    }
                    // set inner html accordingly
                    minuteView.appendChild(docfrag);
                }
            }, {
                key: '_switchToView',
                value: function _switchToView(el) {
                    var me = this;
                    // attach the view change button
                    el.onclick = function () {
                        me._switchToTimeView(me);
                    };
                }
            }, {
                key: '_switchToTimeView',
                value: function _switchToTimeView(me) {
                    var hourView = me._sDialog.hourView,
                        minuteView = me._sDialog.minuteView,
                        hour = me._sDialog.hour,
                        minute = me._sDialog.minute,
                        activeClass = 'mddtp-picker__color--active',
                        hidden = 'mddtp-picker__circularView--hidden',
                        selection = 'mddtp-picker__selection',
                        needle = me._sDialog.needle,
                        circularHolder = me._sDialog.circularHolder,
                        circle = me._sDialog.circle,
                        fakeNeedle = me._sDialog.fakeNeedle,
                        spoke = 60,
                        value = void 0;

                    // toggle view classes
                    hourView.classList.toggle(hidden);
                    minuteView.classList.toggle(hidden);
                    hour.classList.toggle(activeClass);
                    minute.classList.toggle(activeClass);
                    // move the needle to correct position
                    needle.className = '';
                    needle.classList.add(selection);
                    if (mdTimePicker.dialog.view) {
                        value = me._sDialog.sDate.format('m');
                        // move the fakeNeedle to correct position
                        setTimeout(function () {
                            var hOffset = circularHolder.getBoundingClientRect(),
                                cOffset = circle.getBoundingClientRect();

                            fakeNeedle.setAttribute('style', 'left:' + (cOffset.left - hOffset.left) + 'px;top:' + (cOffset.top - hOffset.top) + 'px');
                        }, 300);
                    } else {
                        spoke = 24;
                        value = parseInt(me._sDialog.sDate.format('H'), 10);
                        // CHANGED exception for 24 => 0 issue #58
                        if (value === 0) {
                            value = 24;
                        }
                    }
                    var rotationClass = me._calcRotation(spoke, parseInt(value, 10));
                    if (rotationClass) {
                        needle.classList.add(rotationClass);
                    }
                    // toggle the view type
                    mdTimePicker.dialog.view = !mdTimePicker.dialog.view;
                }
            }, {
                key: '_addClockEvent',
                value: function _addClockEvent() {
                    var me = this,
                        hourView = this._sDialog.hourView,
                        minuteView = this._sDialog.minuteView,
                        sClass = 'mddtp-picker__cell--selected';

                    hourView.onclick = function (e) {
                        var sHour = 'mddtp-hour__selected',
                            selectedHour = document.getElementById(sHour),
                            setHour = 0;

                        if (e.target && e.target.nodeName == 'SPAN') {
                            // clear the previously selected hour
                            selectedHour.id = '';
                            selectedHour.classList.remove(sClass);
                            // select the new hour
                            e.target.parentNode.classList.add(sClass);
                            e.target.parentNode.id = sHour;
                            // set the sDate according to 24 or 12 hour mode
                            setHour = parseInt(e.target.textContent, 10);
                            me._sDialog.sDate.hour(setHour);
                            // set the display hour
                            me._sDialog.hour.textContent = e.target.textContent;
                            // switch the view
                            me._switchToTimeView(me);
                        }
                    };
                    minuteView.onclick = function (e) {
                        var sMinute = 'mddtp-minute__selected',
                            selectedMinute = document.getElementById(sMinute),
                            setMinute = 0;

                        if (e.target && e.target.nodeName == 'SPAN') {
                            // clear the previously selected hour
                            if (selectedMinute) {
                                selectedMinute.id = '';
                                selectedMinute.classList.remove(sClass);
                            }
                            // select the new minute
                            e.target.parentNode.classList.add(sClass);
                            e.target.parentNode.id = sMinute;
                            // set the sDate minute
                            setMinute = e.target.textContent;
                            me._sDialog.sDate.minute(setMinute);
                            // set the display minute
                            me._sDialog.minute.textContent = setMinute;
                            // switch the view
                            me._switchToTimeView(me);
                        }
                    };
                }
            }, {
                key: '_dragDial',
                value: function _dragDial() {
                    var me = this,
                        needle = this._sDialog.needle,
                        circle = this._sDialog.circle,
                        fakeNeedle = this._sDialog.fakeNeedle,
                        circularHolder = this._sDialog.circularHolder,
                        minute = this._sDialog.minute,
                        quick = 'mddtp-picker__selection--quick',
                        selection = 'mddtp-picker__selection',
                        selected = 'mddtp-picker__cell--selected',
                        rotate = 'mddtp-picker__cell--rotate-',
                        hOffset = circularHolder.getBoundingClientRect(),
                        divides = void 0,
                        fakeNeedleDraggabilly = new Draggabilly(fakeNeedle, {
                            containment: !0
                        });

                    fakeNeedleDraggabilly.on('pointerDown', function (e) {
                        console.info('pointerDown', e);
                        hOffset = circularHolder.getBoundingClientRect();
                    });
                    /**
                     * netTrek
                     * fixes for iOS - drag
                     */
                    fakeNeedleDraggabilly.on('pointerMove', function (e) {

                        var clientX = e.clientX,
                            clientY = e.clientY;


                        if (clientX === undefined) {

                            if (e.pageX === undefined) {
                                if (e.touches && e.touches.length > 0) {
                                    clientX = e.touches[0].clientX;
                                    clientY = e.touches[0].clientY;
                                } else {
                                    console.error('coult not detect pageX, pageY');
                                }
                            } else {
                                clientX = pageX - document.body.scrollLeft - document.documentElement.scrollLeft;
                                clientY = pageY - document.body.scrollTop - document.documentElement.scrollTop;
                            }
                        }
                        console.info('Drag clientX', clientX, clientY, e);

                        var xPos = clientX - hOffset.left - hOffset.width / 2,
                            yPos = clientY - hOffset.top - hOffset.height / 2,
                            slope = Math.atan2(-yPos, xPos);

                        needle.className = '';
                        if (slope < 0) {
                            slope += 2 * Math.PI;
                        }
                        slope *= 180 / Math.PI;
                        slope = 360 - slope;
                        if (slope > 270) {
                            slope -= 360;
                        }
                        divides = parseInt(slope / 6);
                        var same = Math.abs(6 * divides - slope),
                            upper = Math.abs(6 * (divides + 1) - slope);

                        if (upper < same) {
                            divides++;
                        }
                        divides += 15;
                        needle.classList.add(selection);
                        needle.classList.add(quick);
                        needle.classList.add(rotate + divides * 2);
                    });
                    /**
                     * netTrek
                     * fixes for iOS - drag
                     */
                    fakeNeedleDraggabilly.on('pointerUp', function (e) {
                        var minuteViewChildren = me._sDialog.minuteView.getElementsByTagName('div'),
                            sMinute = 'mddtp-minute__selected',
                            selectedMinute = document.getElementById(sMinute),
                            cOffset = circle.getBoundingClientRect();

                        fakeNeedle.setAttribute('style', 'left:' + (cOffset.left - hOffset.left) + 'px;top:' + (cOffset.top - hOffset.top) + 'px');
                        needle.classList.remove(quick);
                        var select = divides;
                        if (select === 1) {
                            select = 60;
                        }
                        select = me._nearestDivisor(select, 5);
                        // normalize 60 => 0
                        if (divides === 60) {
                            divides = 0;
                        }
                        // remove previously selected value
                        if (selectedMinute) {
                            selectedMinute.id = '';
                            selectedMinute.classList.remove(selected);
                        }
                        // add the new selected
                        if (select > 0) {
                            select /= 5;
                            select--;
                            minuteViewChildren[select].id = sMinute;
                            minuteViewChildren[select].classList.add(selected);
                        }
                        minute.textContent = me._numWithZero(divides);
                        me._sDialog.sDate.minutes(divides);
                    });
                }
            }, {
                key: '_attachEventHandlers',
                value: function _attachEventHandlers() {
                    var me = this,
                        ok = this._sDialog.ok,
                        cancel = this._sDialog.cancel;

                    cancel.onclick = function () {
                        me.toggle();
                        if (me.onCancel) {
                            me.onCancel();
                        }
                    };
                    ok.onclick = function () {
                        me._init = me._sDialog.sDate;
                        me.toggle();
                        if (me.onOK) {
                            me.onOK(me._init);
                        }
                    };
                }
            }, {
                key: '_setButtonText',
                value: function _setButtonText() {
                    this._sDialog.cancel.textContent = this._cancel;
                    this._sDialog.ok.textContent = this._ok;
                }
            }, {
                key: '_nearestDivisor',
                value: function _nearestDivisor(number, divided) {
                    if (number % divided === 0) {
                        return number;
                    } else if ((number - 1) % divided === 0) {
                        return number - 1;
                    } else if ((number + 1) % divided === 0) {
                        return number + 1;
                    }
                    return -1;
                }
            }, {
                key: '_numWithZero',
                value: function _numWithZero(n) {
                    return n > 9 ? '' + n : '0' + n;
                }
            }, {
                key: '_fillText',
                value: function _fillText(el, text) {
                    if (el.firstChild) {
                        el.firstChild.nodeValue = text;
                    } else {
                        el.appendChild(document.createTextNode(text));
                    }
                }
            }, {
                key: '_addId',
                value: function _addId(el, id) {
                    el.id = 'mddtp-' + this._type + '__' + id;
                }
            }, {
                key: '_addClass',
                value: function _addClass(el, aClass, more) {
                    el.classList.add('mddtp-picker__' + aClass);
                    var i = 0;
                    if (more) {
                        i = more.length;
                        more.reverse();
                    }
                    while (i--) {
                        el.classList.add(more[i]);
                    }
                }
            }, {
                key: '_calcRotation',
                value: function _calcRotation(spoke, value) {
                    var start = spoke / 12 * 3;
                    // set clocks top and right side value
                    if (spoke === 12) {
                        value *= 10;
                    } else if (spoke === 24) {
                        value *= 5;
                    } else {
                        value *= 2;
                    }
                    // special case for 00 => 60
                    if (spoke === 60 && value === 0) {
                        value = 120;
                    }
                    return 'mddtp-picker__cell--rotate-' + value;
                }
            }], [{
                key: 'dialog',
                get: function get() {
                    return _dialog;
                },
                set: function set(value) {
                    mdTimePicker.dialog = value;
                }
            }]);

            return mdTimePicker;
        }();

    exports.default = mdTimePicker;
});