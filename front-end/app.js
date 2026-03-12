var app = angular.module('CalculatorApp', []);
app.controller('CalcController', function($scope, $http) {
    $scope.doCalc = function(op) {
        $http.post('/api/' + op, { number: $scope.val, power: $scope.pow })
            .then(function(res) {
                $scope.result = res.data.result;
            });
    };
});