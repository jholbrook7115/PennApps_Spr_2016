var websiteControllers = angular.module('websiteControllers', ['ngMaterial']);



websiteControllers.controller('homeController', function($scope){

});

websiteControllers.controller('activeShowsController', function($scope, $http){
	$scope.ActiveShows = function getActiveShows(){
	$http({
		method: 'GET',
		url: '/series/series_title',
	}).then(function success(response){
		return response.data;
	}, function errorCallback(response){
		console.log("Error: " + response);
		return;
	});
}
});

websiteControllers.controller('pastShowsController', function($scope, $http){
	$scope.PastShows = function pastShows(){

		return { name: 'This is a placeholder. Feature Not Functional Yet',
				link: '',
				image: ''};
	// 	$http({
	// 	method: 'GET',
	// 	url: '/api/series_title',
	// }).then(function success(response){
	// 	for(var i =0; i < response.length; i++){
	// 		//build the PastShows JSON Object here
	// 		return { name: 'This is a placeholder. Feature Not Functional Yet',
	// 			link: '',
	// 			image: ''};
				
	// 	}
		
	// }, function errorCallback(response){
	// 	console.log("Error: " + response);
	// 	return;
	// });
	}
});

websiteControllers.controller('chatLogController', function($scope){
	$scope.chatLog = [{
		id: '00000001',
		socket1: 'user1',
		content: 'example content 1'
	},{
		id: '00000002',
		socket1: 'user2',
		content: 'example content 2'
	} ];

	// temp http request for updating the chatlog.  Run this every few seconds or something
	// will be replaced by a web socket 
	// $http({
	// 	method: 'GET',
	// 	url: '/chat'
	// }).then(function success(response){
	// 	//get the chatlog
	// }, function errorCallback(response){
	// 	console.log("Error: " + response);
	// });


});
