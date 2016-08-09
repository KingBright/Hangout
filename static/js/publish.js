var module = angular.module("Publish",
    [
        "ngMaterial",
        "ngAnimate",
        "ngAria",
        "ngMessages"
    ]
    , function ($interpolateProvider, $mdThemingProvider) {
        $interpolateProvider.startSymbol('[[');
        $interpolateProvider.endSymbol(']]');
        $mdThemingProvider.theme('default');
    })

module.controller("PublishController", function ($scope) {
    $scope.freeMode = false;
    $scope.hasPunishment = false;

    $scope.startDate
    $scope.endDate
    $scope.date
    $scope.time

    $scope.startDateFormatted
    $scope.endDateFormatted
    $scope.dateFormatted

    $scope.payer = 0
    $scope.payers = [{label: "AA", value: 0}, {label: "我自己", value: 1}]

    $scope.current = new Date();

    $scope.$watch('startDate', function (newValue, oldValue) {
        $scope.startDateFormatted = moment(newValue).format("YYYY-MM-DD")
    });
    $scope.$watch('endDate', function (newValue, oldValue) {
        $scope.endDateFormatted = moment(newValue).format("YYYY-MM-DD")
    });
    $scope.$watch('date', function (newValue, oldValue) {
        $scope.dateFormatted = moment(newValue).format("YYYY-MM-DD")
    });

    $scope.dateFilter = function (date) {
        return moment(date).isAfter(moment(this.current), 'day') || moment(date).isSame(moment(this.current), 'day')
    };

    $scope.startDateFilter = function (date) {
        var target = moment(date)
        if ($scope.endDate) {
            return target.isBefore(moment($scope.endDate), 'day') && (target.isAfter(moment(this.current), 'day') || target.isSame(moment(this.current), 'day'))
        }
        return target.isAfter(moment(this.current), 'day') || target.isSame(moment(this.current), 'day')
    };

    $scope.endDateFilter = function (date) {
        var target = moment(date)
        if ($scope.startDate) {
            return target.isAfter(moment($scope.startDate), 'day')
        }
        return target.isAfter(moment(this.current), 'day')
    };

    var dialog = new mdTimePicker.default();
    dialog.onOK = function (date) {
        $scope.time = moment(date).format("HH:mm")
    }

    $scope.showClock = function () {
        dialog.toggle()
    }
})