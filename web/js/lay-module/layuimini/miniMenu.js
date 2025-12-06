/**
 * date:2020/02/27
 * author:Mr.Chung
 * version:2.0
 * description:layuimini 菜单框架扩展
 */
layui.define(["element","laytpl" ,"jquery"], function (exports) {
    var element = layui.element,
        $ = layui.$,
        laytpl = layui.laytpl,
        layer = layui.layer;

    var miniMenu = {

        /**
         * 菜单初始化
         * @param options.menuList   菜单数据信息
         * @param options.multiModule 是否开启多模块
         * @param options.menuChildOpen 是否展开子菜单
         */
        render: function (options) {
            options.menuList = options.menuList || [];
            options.multiModule = options.multiModule || false;
            options.menuChildOpen = options.menuChildOpen || false;
            if (options.multiModule) {
                miniMenu.renderMultiModule(options.menuList, options.menuChildOpen);
            } else {
                miniMenu.renderSingleModule(options.menuList, options.menuChildOpen);
            }
            miniMenu.listen();
        },

        /**
         * 单模块
         * @param menuList 菜单数据
         * @param menuChildOpen 是否默认展开
         */
        renderSingleModule: function (menuList, menuChildOpen) {
            menuList = menuList || [];
            var leftMenuHtml = '',
                childOpenClass = '',
                leftMenuCheckDefault = 'layui-this';
            var me = this ;
            if (menuChildOpen) childOpenClass = ' layui-nav-itemed';
            leftMenuHtml = this.renderLeftMenu(menuList,{ childOpenClass:childOpenClass }) ;
            $('.layui-layout-body').addClass('layuimini-single-module'); //单模块标识
            $('.layuimini-header-menu').remove();
            // 有时 laytpl 渲染嵌套 children 时会被转义为 &lt;li&gt; 字符串，做防御性解码
            try{
                var _html = leftMenuHtml;
                // 如果被转义（可能是单次或多次转义），反复解码直到不含实体或达到最大迭代次数
                if (typeof _html === 'string' && _html.indexOf('&lt;') !== -1) {
                    var attempts = 0, prev;
                    while (attempts < 5 && _html.indexOf('&lt;') !== -1) {
                        prev = _html;
                        _html = $('<div/>').html(_html).text();
                        if (_html === prev) break;
                        attempts++;
                    }
                }
                $('.layuimini-menu-left').html(_html);
                // 确保渲染并显示第一个菜单
                element.init();
                var $left = $('.layuimini-menu-left');
                $left.find('ul').removeClass('layui-hide');
                $left.find('ul').first().addClass('layui-this').removeClass('layui-hide');
                element.render();
            } catch (e) {
                // 回退：直接插入原始字符串（尽管可能显示为文本）
                // 解码失败时不打印冗余日志，直接回退插入
                $('.layuimini-menu-left').html(leftMenuHtml);
            }

            // 确保渲染并显示第一个菜单
            element.init();
            try {
                var $left = $('.layuimini-menu-left');
                // 移除多余的隐藏类，显示第一个菜单
                $left.find('ul').removeClass('layui-hide');
                $left.find('ul').first().addClass('layui-this').removeClass('layui-hide');
                element.render();
            } catch (e) {
                console.warn('renderSingleModule post-render adjust failed', e);
            }
        },

        /**
         * 渲染一级菜单
         */
        compileMenu: function(menu,isSub){
            // 使用非转义输出 {{= d.children }} 以兼容 laytpl 在新版中对变量的转义行为
            var menuHtml = '<li {{#if( d.menu){ }}  data-menu="{{d.menu}}" {{#}}} class="layui-nav-item menu-li {{d.childOpenClass}} {{d.className}}"  {{#if( d.id){ }}  id="{{d.id}}" {{#}}}> <a {{#if( d.href){ }} layuimini-href="{{d.href}}" {{#}}} {{#if( d.target){ }}  target="{{d.target}}" {{#}}} href="javascript:;">{{#if( d.icon){ }}  <i class="{{d.icon}}"></i> {{#}}} <span class="layui-left-nav">{{d.title}}</span></a>  {{# if(d.children){}} {{= d.children }} {{#}}} </li>' ;
            if(isSub){
                // dd 子模板中同样使用非转义输出以插入子节点 HTML
                menuHtml = '<dd class="menu-dd {{d.childOpenClass}} {{ d.className }}"> <a href="javascript:;"  {{#if( d.menu){ }}  data-menu="{{d.menu}}" {{#}}} {{#if( d.id){ }}  id="{{d.id}}" {{#}}} {{#if(( !d.child || !d.child.length ) && d.href){ }} layuimini-href="{{d.href}}" {{#}}} {{#if( d.target){ }}  target="{{d.target}}" {{#}}}> {{#if( d.icon){ }}  <i class="{{d.icon}}"></i> {{#}}} <span class="layui-left-nav"> {{d.title}}</span></a> {{# if(d.children){}} {{= d.children }} {{#}}}</dd>'
            }
            return laytpl(menuHtml).render(menu);
        },
        compileMenuContainer :function(menu,isSub){
            // 父容器使用非转义输出 children
            var wrapperHtml = '<ul class="layui-nav layui-nav-tree layui-left-nav-tree {{d.className}}" id="{{d.id}}">{{= d.children }}</ul>' ;
            if(isSub){
                wrapperHtml = '<dl class="layui-nav-child ">{{= d.children }}</dl>' ;
            }
            if(!menu.children){
                return "";
            }
            return laytpl(wrapperHtml).render(menu);
        },

        each:function(list,callback){
            var _list = [];
            for(var i = 0 ,length = list.length ; i<length ;i++ ){
                _list[i] = callback(i,list[i]) ;
            }
            return _list ;
        },
        renderChildrenMenu:function(menuList,options){
            var me = this ;
            menuList = menuList || [] ;
            var html = this.each(menuList,function (idx,menu) {
                if(menu.child && menu.child.length){
                    menu.children = me.renderChildrenMenu(menu.child,{ childOpenClass: options.childOpenClass || '' });
                }
                menu.className = "" ;
                menu.childOpenClass = options.childOpenClass || ''
                return me.compileMenu(menu,true)
            }).join("");
            return me.compileMenuContainer({ children:html },true)
        },
        renderLeftMenu :function(leftMenus,options){
            options = options || {};
            var me = this ;
            var leftMenusHtml =  me.each(leftMenus || [],function (idx,leftMenu) { // 左侧菜单遍历
                var children = me.renderChildrenMenu(leftMenu.child, { childOpenClass:options.childOpenClass });
                var leftMenuHtml = me.compileMenu({
                    href: leftMenu.href,
                    target: leftMenu.target,
                    childOpenClass: options.childOpenClass,
                    icon: leftMenu.icon,
                    title: leftMenu.title,
                    children: children,
                    className: '',
                });
                return leftMenuHtml ;
            }).join("");

            leftMenusHtml = me.compileMenuContainer({ id:options.parentMenuId,className:options.leftMenuCheckDefault,children:leftMenusHtml }) ;
            return leftMenusHtml ;
        },
        /**
         * 多模块
         * @param menuList 菜单数据
         * @param menuChildOpen 是否默认展开
         */
        renderMultiModule: function (menuList, menuChildOpen) {
            menuList = menuList || [];
            var me = this ;
            var headerMenuHtml = '',
                headerMobileMenuHtml = '',
                leftMenuHtml = '',
                leftMenuCheckDefault = 'layui-this',
                childOpenClass = '',
                headerMenuCheckDefault = 'layui-this';

            if (menuChildOpen) childOpenClass = ' layui-nav-itemed';
            var headerMenuHtml = this.each(menuList, function (index, val) { //顶部菜单渲染
                var menu = 'multi_module_' + index ;
                var id = menu+"HeaderId";
                var topMenuItemHtml = "" ;
                topMenuItemHtml = me.compileMenu({
                    className:headerMenuCheckDefault,
                    menu:menu,
                    id:id,
                    title:val.title,
                    href:"",
                    target:"",
                    children:""
                });
                leftMenuHtml+=me.renderLeftMenu(val.child,{
                    parentMenuId:menu,
                    childOpenClass:childOpenClass,
                    leftMenuCheckDefault:leftMenuCheckDefault
                });
                headerMobileMenuHtml +=me.compileMenu({ id:id,menu:menu,id:id,icon:val.icon, title:val.title, },true);
                headerMenuCheckDefault = "";
                leftMenuCheckDefault = "layui-hide" ;
                return topMenuItemHtml ;
            }).join("");
            $('.layui-layout-body').addClass('layuimini-multi-module'); //多模块标识
            // header/menu 也可能被 laytpl 转义，统一做防御性解码再插入
            try{
                function decodeEntities(str) {
                    if (typeof str !== 'string') return str;
                    var attempts = 0, prev;
                    while (attempts < 5 && str.indexOf('&lt;') !== -1) {
                        prev = str;
                        str = $('<div/>').html(str).text();
                        if (str === prev) break;
                        attempts++;
                    }
                    return str;
                }

                var _header = decodeEntities(headerMenuHtml);
                $('.layuimini-menu-header-pc').html(_header); //电脑

                var _left = decodeEntities(leftMenuHtml);
                $('.layuimini-menu-left').html(_left);

                // 处理当 HTML 被当作文本节点插入（例如页面显示 "<li ...>" 文本）的情况
                try{
                    var $uls = $('.layuimini-menu-left').find('ul, dl');
                    $uls.each(function(i, el){
                        try{
                            // 如果元素没有子元素但文本内容包含 <li 或 <dd，则把文本作为 HTML 解析一次
                            if (el.children.length === 0) {
                                var txt = el.textContent || el.innerText || '';
                                if (txt && txt.indexOf('<li') !== -1) {
                                    el.innerHTML = txt;
                                } else if (txt && txt.indexOf('<dd') !== -1) {
                                    el.innerHTML = txt;
                                }
                            }
                        }catch(e){/* ignore per-item errors */}
                    });
                }catch(e){/* ignore post-parse conversion errors silently */}
                var _mobile = decodeEntities(headerMobileMenuHtml);
                $('.layuimini-menu-header-mobile').html(_mobile); //手机

                // 确保左侧菜单显示并渲染
                element.init();
                var $left = $('.layuimini-menu-left');
                if ($left.find('ul.layui-hide').length === $left.find('ul').length) {
                    $left.find('ul').removeClass('layui-hide');
                }
                $left.find('ul').first().addClass('layui-this').removeClass('layui-hide');
                element.render();
            } catch (e) {
                // 解码/插入 header/menu 失败时不打印调试日志，直接回退原始内容
                $('.layuimini-menu-header-pc').html(headerMenuHtml);
                $('.layuimini-menu-left').html(leftMenuHtml);
                $('.layuimini-menu-header-mobile').html(headerMobileMenuHtml);
            }
            // 初始化 element 并确保左侧第一个子菜单显示
            element.init();
            try {
                var $left = $('.layuimini-menu-left');
                // 如果菜单默认都被设置为 layui-hide，则显示第一个
                if ($left.find('ul.layui-hide').length === $left.find('ul').length) {
                    $left.find('ul').removeClass('layui-hide');
                }
                // 给第一个左侧菜单标记为激活
                $left.find('ul').first().addClass('layui-this').removeClass('layui-hide');
                // 渲染 element
                element.render();
            } catch (e) {
                console.warn('renderMultiModule post-render adjust failed', e);
            }
        },

        /**
         * 监听
         */
        listen: function () {

            /**
             * 菜单模块切换
             */
            $('body').on('click', '[data-menu]', function () {
                var loading = layer.load(0, {shade: false, time: 2 * 1000});
                var menuId = $(this).attr('data-menu');
                // header
                $(".layuimini-header-menu .layui-nav-item.layui-this").removeClass('layui-this');
                $(this).addClass('layui-this');
                // left
                $(".layuimini-menu-left .layui-nav.layui-nav-tree.layui-this").addClass('layui-hide');
                $(".layuimini-menu-left .layui-nav.layui-nav-tree.layui-this.layui-hide").removeClass('layui-this');
                $("#" + menuId).removeClass('layui-hide');
                $("#" + menuId).addClass('layui-this');
                layer.close(loading);
            });

            /**
             * 菜单缩放
             */
            $('body').on('click', '.layuimini-site-mobile', function () {
                var loading = layer.load(0, {shade: false, time: 2 * 1000});
                var isShow = $('.layuimini-tool [data-side-fold]').attr('data-side-fold');
                if (isShow == 1) { // 缩放
                    $('.layuimini-tool [data-side-fold]').attr('data-side-fold', 0);
                    $('.layuimini-tool [data-side-fold]').attr('class', 'fa fa-indent');
                    $('.layui-layout-body').removeClass('layuimini-all');
                    $('.layui-layout-body').addClass('layuimini-mini');
                } else { // 正常
                    $('.layuimini-tool [data-side-fold]').attr('data-side-fold', 1);
                    $('.layuimini-tool [data-side-fold]').attr('class', 'fa fa-outdent');
                    $('.layui-layout-body').removeClass('layuimini-mini');
                    $('.layui-layout-body').addClass('layuimini-all');
                    layer.close(window.openTips);
                }
                element.init();
                layer.close(loading);
            });
            /**
             * 菜单缩放
             */
            $('body').on('click', '[data-side-fold]', function () {
                var loading = layer.load(0, {shade: false, time: 2 * 1000});
                var isShow = $('.layuimini-tool [data-side-fold]').attr('data-side-fold');
                if (isShow == 1) { // 缩放
                    $('.layuimini-tool [data-side-fold]').attr('data-side-fold', 0);
                    $('.layuimini-tool [data-side-fold]').attr('class', 'fa fa-indent');
                    $('.layui-layout-body').removeClass('layuimini-all');
                    $('.layui-layout-body').addClass('layuimini-mini');
                    // $(".menu-li").each(function (idx,el) {
                    //     $(el).addClass("hidden-sub-menu");
                    // });

                } else { // 正常
                    $('.layuimini-tool [data-side-fold]').attr('data-side-fold', 1);
                    $('.layuimini-tool [data-side-fold]').attr('class', 'fa fa-outdent');
                    $('.layui-layout-body').removeClass('layuimini-mini');
                    $('.layui-layout-body').addClass('layuimini-all');
                    // $(".menu-li").each(function (idx,el) {
                    //     $(el).removeClass("hidden-sub-menu");
                    // });
                    layer.close(window.openTips);
                }
                element.init();
                layer.close(loading);
            });

            /**
             * 手机端点开模块
             */
            $('body').on('click', '.layuimini-header-menu.layuimini-mobile-show dd', function () {
                var loading = layer.load(0, {shade: false, time: 2 * 1000});
                var check = $('.layuimini-tool [data-side-fold]').attr('data-side-fold');
                if(check === "1"){
                    $('.layuimini-site-mobile').trigger("click");
                    element.init();
                }
                layer.close(loading);
            });
        },

    };


    exports("miniMenu", miniMenu);
});
