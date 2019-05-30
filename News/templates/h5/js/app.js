angular.module('myApp', ['ionic', 'controllers', 'ui.router'])
	.run(function($ionicPlatform,$rootScope,$location) {
     	$rootScope.$on('$locationChangeStart', function () {
	        try{
	          sxkui.baidu($location.path());
	        }catch(e){
	        }
      	});
		$ionicPlatform.ready(function() {
			if(window.cordova && window.cordova.plugins && window.cordova.plugins.Keyboard) {
				cordova.plugins.Keyboard.hideKeyboardAccessoryBar(true);
				cordova.plugins.Keyboard.disableScroll(true);
			}
			if(window.StatusBar) {
				StatusBar.styleDefault();
			}
		});
	})

.config(function($stateProvider, $urlRouterProvider, $ionicConfigProvider) {

	$ionicConfigProvider.platform.ios.tabs.style('standard');
	$ionicConfigProvider.platform.ios.tabs.position('bottom');
	$ionicConfigProvider.platform.android.tabs.style('standard');
	$ionicConfigProvider.platform.android.tabs.position('standard');
	$ionicConfigProvider.platform.ios.navBar.alignTitle('center');
	$ionicConfigProvider.platform.android.navBar.alignTitle('center');
	$ionicConfigProvider.platform.ios.backButton.previousTitleText('').icon('ion-ios-arrow-thin-left');
	$ionicConfigProvider.platform.android.backButton.previousTitleText('').icon('ion-android-arrow-back');
	$ionicConfigProvider.platform.ios.views.transition('ios');
	$ionicConfigProvider.platform.android.views.transition('android');

	// $urlRouterProvider.otherwise('/tab/home');

	$stateProvider.state('tab', {
		url: '/tab',
		templateUrl: 'templates/tabs.html',
		controller: 'tabCtrl'
	});
	$stateProvider.state('tab.home', {
		url: '/home',
		cache:'false',
		views: {
			"home": {
				templateUrl: 'templates/home.html',
				controller: 'homeCtrl'
			}
		}
	});
	$stateProvider.state('tab.order', {
		url: '/order',
		views: {
			"order": {
				templateUrl: 'templates/order.html'
					//      controller:'homeCtrl'
			}
		}
	});
	$stateProvider.state('tab.my', {
		url: '/my',
		cache:'false',
		views: {
			"my": {
				templateUrl: 'templates/my.html',
				controller: 'myCtrl'
			}
		}
	});
	$stateProvider.state('service', {
		url: '/service',
		templateUrl: 'templates/service.html',
		controller: 'serviceCtrl'
	});
	$stateProvider.state('edit', {
		url: '/edit',
		templateUrl: 'templates/edit.html',
		controller: 'editCtrl'
	});
	$stateProvider.state('new_message', {
		url: '/new_message',
		templateUrl: 'templates/new_message.html',
		controller: 'new_messageCtrl'
	});
	$stateProvider.state('searchCard', {
		url: '/searchCard',
		templateUrl: 'templates/searchCard.html',
		controller: 'searchCardCtrl'
	});
	$stateProvider.state('message_list', {
		url: '/message_list',
		cache: 'false',
		templateUrl: 'templates/message/messageList.html',
		controller: 'message_listCtrl'
	});
		$stateProvider.state('message_list_detail', {
		url: '/message_list_detail',
		templateUrl: 'templates/message/messageListDetail.html',
		controller: 'message_list_detailCtrl'
	});

});