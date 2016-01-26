var app = angular.module( 'RootApp', [ 'ngMaterial', 'ngRoute', 'websiteControllers' ] );

	app.config(function($routeProvider){
	$routeProvider
		.when('/', {
			templateUrl: 'templates/pages/home.html',
			controller: 'homeController'
		})
		.when('/active-shows', {
			templateUrl: 'templates/pages/activeShows.html',
			controller: 'activeShowsController'
		})
		.when('/past-shows', {
			templateUrl: 'templates/pages/pastShows.html',
			controller: 'pastShowsController'
		})
		.when('/chat-room-list', {
			templateUrl: 'templates/pages/chatRoom.html',
			controller: 'chatLogController'
		})
		.otherwise({
			redirectTo: '/'
		});
});

app.controller('MainCtrl', function($scope, $mdSidenav){

	$scope.sidebarObjects = [{
		name: 'Home',
		link: '#home',
		description: 'The home page'

	},{
		name: 'Live Now',
		link: '#active-shows',
		description: 'Shows airing right now'
	},{
		name: 'Past Shows',
		link: '#past-shows',
		description: 'Shows that you may have missed'
	},{
		name: 'Chat Rooms',
		link: '#chat-room-list',
		description: 'A list of all the chat rooms'
	}];


	$scope.tmpUserJSON =[{
	name: 'user1',
	color: 'red'
	},{
	name: 'user2',
	color: 'pink'
	},{
	name: 'user3',
	color: 'purple'
	}];

	$scope.openLeftMenu = function(){
		$mdSidenav('left').toggle();
	};
});