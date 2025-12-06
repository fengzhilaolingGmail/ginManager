/**
 * Layui v2.6.3 到 v2.13.2 兼容性修复
 * 解决升级过程中可能出现的兼容性问题
 */

// 兼容性检查
layui.use(['jquery', 'layer'], function(){
    var $ = layui.jquery;
    var layer = layui.layer;
    
    // compatibility script loaded (debug logs removed)
    
    // 检查版本
    if (layui.v && layui.v.startsWith('2.13')) {
    // Layui version detected
        
        // 修复可能存在的配置问题
        if (layui.cache && layui.cache.version) {
            // cache version available
        }
        
        // 确保jQuery正确加载
        if (typeof $ === 'undefined') {
            console.error('jQuery未正确加载，尝试重新加载...');
            // 这里可以添加jQuery的备用加载逻辑
        }
        
        // 版本升级提示
        if (console && console.info) {
            // upgrade info (suppressed)
        }
        
    } else {
        // version check abnormal (suppressed)
    }
});

// 全局错误处理
window.addEventListener('error', function(e) {
    if (e.filename && e.filename.includes('layui')) {
        // suppressed detailed error logging for layui files
    }
});

// 兼容性修复函数
window.layuiCompatFix = {
    // 修复旧版本API调用
    fixOldApiCalls: function() {
        // 如果有旧版本的API调用，在这里添加修复逻辑
    // check old API calls (suppressed log)
    },
    
    // 验证关键组件
    validateComponents: function() {
        var components = ['form', 'table', 'layer', 'element', 'jquery'];
        var results = {};
        
        layui.use(components, function() {
            for (var i = 0; i < components.length; i++) {
                var compName = components[i];
                try {
                    var comp = layui[compName];
                    results[compName] = comp ? 'ok' : 'missing';
                } catch (e) {
                    results[compName] = 'error: ' + e.message;
                }
            }
            
            // component validation results (suppressed)
            return results;
        });
    },
    
    // 显示兼容性报告
    showCompatReport: function() {
        var report = 'Layui v2.13.2 兼容性报告\\n';
        report += '========================\\n';
        report += '版本: ' + (layui.v || 'unknown') + '\\n';
        report += '状态: ' + (layui.v === '2.13.2' ? '✅ 升级成功' : '⚠️ 版本异常') + '\\n';
        report += '时间: ' + new Date().toLocaleString() + '\\n';
        
        return report;
    }
};

// 页面加载完成后执行兼容性检查
layui.use(['jquery'], function($) {
    $(document).ready(function() {
    // executing compatibility checks (suppressed)
            layuiCompatFix.fixOldApiCalls();
        
        // 延迟执行组件验证，确保所有模块都加载完成
        setTimeout(function() {
            layuiCompatFix.validateComponents();
        }, 1000);
        
        // 输出兼容性报告
        setTimeout(function() {
            // compatibility report suppressed
        }, 2000);
    });
});