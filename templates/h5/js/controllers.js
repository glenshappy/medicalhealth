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
			sxkui.activity('medical.home');
		}
		$scope.goRecord = function(){
			sxkui.activity('user.dangan');
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
		sxkui.json('index/init', function(data) {
			$scope.data = data.datas;
			if($scope.data.weather.warning != null){
				$('#home>.vip_link>.weather_icon').css('display','block');
			}
		});
		sxkui.json('index/tiantiance', function(data) {
			$scope.daylist = data.datas;
		});
		//首屏提醒
		$scope.seewarn = function(){
			$('#home_warn').slideDown();
			$('div.tab-nav.tabs').css('display','none');
		}
		$scope.threeAlertDatas = function(){
			sxkui.json('dialog', function(data) {
			if(data.datas.length != '0'){
				$scope.alertDatas = data.datas;
				$scope.onealertdata = $scope.alertDatas[$scope.alertIndex];
				$('#home_all_alert').css('display','flex');
				$('div.tab-nav.tabs').css('display','none');
			}else{
				$('#home_all_alert').css('display','none');
			}
			});
		}
		
		$scope.warn_slideup = function(){
			$('#home_warn').slideUp();
			$('div.tab-nav.tabs').css('display','flex');
		}
		//首屏提醒接口
		sxkui.json('attention', function(data) {
			$scope.warnMessage = data.datas;
			if($scope.warnMessage.isShow == '1' && $scope.data.user.isActivate == '1'){
				if(sessionStorage.getItem('shift')){
				}else{
					$('#home_warn').slideDown();
						$('div.tab-nav.tabs').css('display','none');
						sessionStorage.setItem('shift','shift');
				}
			}
		});
		//几大弹框
		$scope.alertIndex = 0;
		$scope.threeAlertDatas();
		window.onpageshow = function(event) {　　
			if(event.persisted) {　
				window.location.reload();
			}
		};
		$scope.alertRead = function(name,status,url){
			sxkui.json('dialog/read', {
				'name': name,
				'do':status
			}, function(data) {
				if(data.datas.status == 1){
					if(url){
						location.href = url;
					}
				}
			});
		}
		$scope.alert_cancle = function(name){
			if ($('input#noAlert').is(':checked')) {
                $scope.alertRead(name,'close');
            }else{
        	 	$scope.alertRead(name,'show');
            }
            $scope.alertIndex++;
            if($scope.alertIndex == $scope.alertDatas.length) {
				$('#home_all_alert').css('display','none');
				$('div.tab-nav.tabs').css('display','flex');
			}
            $scope.onealertdata = $scope.alertDatas[$scope.alertIndex];
            $('#home_all_alert>.alert_main>.alert_foot>.noAlert>label').removeClass('default');
            $('input#noAlert').removeAttr("checked");
		}
		$scope.alert_sure = function(url,name){
			if ($('input#noAlert').is(':checked')) {
                $scope.alertRead(name,'close',url);
                $('div.tab-nav.tabs').css('display','flex');
            }else{
            	$scope.alertRead(name,'show',url);
            	$('div.tab-nav.tabs').css('display','flex');
            }
		}
		$scope.noAlert = function(){
		  	if ($(event.target).is(':checked')) {
                $(event.target).next('label').addClass('default');
            } else{
            	 $(event.target).next('label').removeClass('default');
            }
            
		}
		if(localStorage.getItem('guided')){
		}else{
			$('#guide1').css('display','block');
			$('div.tab-nav.tabs').css('display','none');
			localStorage.setItem('guided','guided');
		}
		//新手引导
		$scope.next_guide = function(num){
			if(num == '1'){
				$('#guide2').css('display','block').siblings('div').css('display','none');
			}else if(num == '2'){
				$('#guide3').css('display','block').siblings('div').css('display','none');
			}else if(num == '3'){
				$('#guide3').css('display','none');
				$('div.tab-nav.tabs').css('display','flex');
			}
		}
		//天天康跳转
		$scope.dayhealthdetail = function(url) {
				location.href = url;
			}
			//消息提醒详情页
		$scope.gomessagedetail = function() {
				$state.go('new_message');
			}
			//跳转
		$scope.goDetail1 = function(url) {
			location.href = url;
		}
		$scope.goLink1 = function(url) {
			location.href = url;
		}
		$scope.goListDetail = function(url) {
			location.href = url;
		}
		$scope.goMore = function() {
			sxkui.activity('news.main');
		}
		$scope.newsDetail1 = function(id) {
			sxkui.activity('news.detail', {
				'id': id
			});
		}
		$scope.go_label_list = function(id) {
				window.event ? window.event.cancelBubble = true : e.stopPropagation();
				sxkui.activity('news.labels', {
					'id': id
				});
			}
		//天气详情页
		$scope.iconDetail = function() {
			location.href = $scope.data.weather.url;
		}
		$scope.warnDetail = function(url,queueId){
				var $this = $(event.target);
				if(queueId == 0){
					location.href = url;
					return false;
				}
				sxkui.json('attention/read', {
					queueId: queueId
				}, function(data) {
					if(data.ret == 0){
						if(url == ''){
							$this.parents('.list1').animate({right:"100vw"},800);
							$timeout(function() {
								$this.parents('.list1').css('height','0');
							}, 800);	
						}else{
							location.href = url;
						}
					}
				});
		}
		$scope.del = function(){
			var $this = $(event.target);
			$this.parents('.list1').animate({right:"100vw"},800);
			$timeout(function() {
				$this.parents('.list1').css('height','0');
			}, 800);
		}
		$scope.goJieQi = function(url){
			location.href = url;
		}
		$scope.go_shop_card = function(){
			sxkui.activity('index.show365');
		}
		//首屏提醒展开收起
		$scope.slidedown = function(){
			$(event.target).parents('.middle').css('height','auto');
			$(event.target).parents('.middle').css('max-height','none');
			$(event.target).css('display','none');
			$(event.target).siblings('.close').css('display','block');
		}
		$scope.slideup = function(){
			$(event.target).parents('.middle').css('max-height','174px');
			$(event.target).css('display','none');
			$(event.target).siblings('.open').css('display','block');
		}
		})
	})
	.controller('myCtrl', function($scope, $state, $ionicLoading, $ionicHistory, $stateParams, $ionicPopup) {
		//用户信息
		sxkui.json('user/info', function(data) {
			$scope.data = data.datas;
		});
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