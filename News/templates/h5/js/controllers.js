angular.module('controllers', ["ionic"])
	.controller('indexCtrl', function($scope, $state, $ionicLoading, $ionicHistory, $stateParams, $ionicPopup, $timeout) {
		//加载动画
		$scope.$on('$ionicView.beforeEnter', function() {
			$ionicLoading.show({});
			$timeout(function() {
				$ionicLoading.hide();
			}, 600);
		});
	})
	.controller('tabCtrl', function($scope, $state, $ionicLoading, $ionicHistory, $stateParams, $ionicPopup) {
		$scope.goOut = function() {
			medical_healthui.activity('medical.home');
		}
		$scope.goRecord = function(){
			medical_healthui.activity('user.dangan');
		}
	})
	.controller('homeCtrl', function($scope, $state, $ionicLoading, $ionicHistory, $stateParams, $ionicPopup, $timeout, $ionicSlideBoxDelegate) {
		$(document).ready(function(){
			$scope.slideChanged = function(index) {
				$scope.slideIndex = index;
				if(($ionicSlideBoxDelegate.count() - 1) == index) {
					$timeout(function() {
						$ionicSlideBoxDelegate.slide(0);
					}, 3000);
				}
			}
		//首页数据总接口
			medical_healthui.json('index/index', {method:"get"},function(data) {
			console.log("获取的数据")
			console.log(data)
		});
		})
	})
	.controller('myCtrl', function($scope, $state, $ionicLoading, $ionicHistory, $stateParams, $ionicPopup) {
		//用户信息
		medical_healthui_video.json('user/info', function(data) {
			$scope.data = data.datas;
		});
		$scope.goLogin = function(){
			$state.go('login')
		}
		$scope.goLink1 = function() {
			$state.go('searchCard')
		}
		$scope.goUserModify = function() {
			sxkui.activity('user.modify')
		}
		$scope.gomycard = function() {
			sxkui.activity('user.mycard')
		}
		$scope.report = function() {
			sxkui.activity('news.feedback');
		}
		$scope.messag = function() {
			sxkui.activity('www.message');
		}
		$scope.myorder = function() {
			sxkui.activity('trade.myorder');
		}
		$scope.cart = function() {
//			sxkui.activity('trade.cart');
			sxkui.activity('medical.myservice');
		}
		$scope.goFangLists = function() {
			sxkui.activity('clock.main');
		}
		$scope.help = function() {
			sxkui.activity('news.help');
		}
		$scope.goshop = function() {
			$state.go('service');
		}
		$scope.edit_new = function() {
			$state.go('edit');
		}
		$scope.go_shop_card = function(){
			sxkui.activity('index.show365');
		}
		$scope.go_showcard = function(){
			sxkui.activity('index.show365');
		}
		//系统消息
		$scope.messageList = function(){
			$state.go('message_list');
		}
			//未读消息 
		$scope.unread = function(){
			sxkui.json('message/unread', function(data) {
				$scope.messNum = data.datas.count;
				//$scope.$apply();
			});
		}
		$scope.unread();
		window.onpageshow = function(event) {　　
			if(event.persisted) {　　　　
				$scope.unread();　
			}
		};
		$scope.goaddcard = function(){
			sxkui.activity('user.addcard');
		}
	})
	.controller('serviceCtrl', function($scope, $state, $ionicLoading, $ionicHistory, $stateParams, $ionicPopup) {
		sxkui.json('card/goods', function(data) {
			$scope.data = data.datas;
		});
		$scope.shopDetail = function(url) {
			location.href = url;
		}
	})
	.controller('searchCardCtrl', function($scope, $state, $ionicLoading, $ionicHistory, $stateParams, $ionicPopup) {
		sxkui.setTitle("健康卡查询");
		$scope.search = function() {
			if($('input').val().length == '6') {
				sxkui.json('card/query', {
					cardCode: $('input').val()
				}, function(data) {
					$scope.data = data.datas;
					if($scope.data.status == '不存在') {
						$('#search_card>div>.data').css('display', 'none')
						$('#search_card>div>p').css('display', 'block')
					} else {
						$('#search_card>div>.data').css('display', 'block')
						$('#search_card>div>p').css('display', 'none')
					}
				});
			} else {
				$('.alertOther').fadeIn('slow');
				setTimeout(function() {
					$('.alertOther').fadeOut("slow");
				}, 1000);
				return false;
			}
		}
		$scope.goLiving = function(){
			sxkui.activity('user.addcard');
		}
	})
	.controller('new_messageCtrl', function($scope, $state, $ionicLoading, $ionicHistory, $stateParams, $ionicPopup) {
		sxkui.json('guide', {
			isremind: '1'
		}, function(data) {
			$scope.textList = data.datas;
		});
		$scope.go_detail = function(id, url) {
			sxkui.json('guide/read', {
				queueId: id
			}, function(data) {
				if(data.datas.status == '1') {
					location.href = url;
				}
			});
		}

	})
	.controller('editCtrl', function($scope, $state, $ionicLoading, $ionicHistory, $stateParams, $ionicPopup) {
		$scope.location = function() {
			sxkui.activity('trade.location');
		}
	})
	.controller('message_listCtrl', function($scope, $state, $ionicLoading, $ionicHistory, $stateParams, $ionicPopup,$ionicScrollDelegate) {
		 $scope.data = {
            showDelete: false
        };
		pageIndex = 1;
        $scope.messageLists = [];
    	$scope.messageData = function(){
			sxkui.json('message/lists', {
				type: '3',
				size:'10',
				page:pageIndex++
			}, function(data) {
				if(data.datas.length == '0') {
					$scope.moredata = false;
					return false;
				}else{
					$scope.moredata = true;
				}
				$scope.messageLists = $scope.messageLists.concat(data.datas);
				sxkui.setTitle('消息');
				$scope.$broadcast('scroll.infiniteScrollComplete');
				$ionicScrollDelegate.resize();
			});
    	}
    	$scope.messageData();
    	window.onpageshow = function(event) {　　
			if(event.persisted) {　　　　
				$scope.messageData();
				//$scope.$apply();
			}
		};
        
		//上拉加载更多
		$scope.loadMore = function() {
			$scope.messageData();
		};
		//去详情页
		$scope.goDetail = function(id){
			location.href= '?id=' + id + '#/message_list_detail';
			$(event.target).parents('.messageList').find('span').remove();
		}
        //删除消息
        $scope.onItemDelete = function (id, index) {
            var $this = $(event.target);
            $ionicPopup.confirm({
                template: "<div style='text-align:center;font-family: 微软雅黑;'>确定删除该消息吗？</div>",
                buttons: [{
                    text: "<div class='messagePopup messagePopup1'>取消</div>",
                    onTap: function (e) {
                        $('.item-content').attr('style', '');
                    }
                }, {
                    text: '<div class="messagePopup messagePopup2">确定</div>',
                    onTap: function (e) {
                        $ionicLoading.show({});
                    	sxkui.json('message/delete', {
							id:id
						}, function(data) {
							if(data.datas.status == '1'){
                                $this.parents('.messageList').remove();
							    $ionicLoading.hide();
							}
						});	
                    }
                }]
            });
        };
	})
	.controller('message_list_detailCtrl', function($scope, $state, $ionicLoading, $ionicHistory, $stateParams, $ionicPopup) {
		sxkui.json('message/detail', {
			id:sxkui.getQuery('id')
		}, function(data) {
			$scope.content = data.datas;
		 	sxkui.setTitle($scope.content.title);
		 	$('#messageListDetail>.summary').html($scope.content.content);
		});	
	})