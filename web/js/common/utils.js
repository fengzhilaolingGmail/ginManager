;(function(){
    // 公共工具：将 UpdatedAt 转换为更易读的时间格式（YYYY-MM-DD HH:mm:ss），并附带本地化时区显示（例如：亚洲/上海）
    function pad(n){ return n < 10 ? '0' + n : n }

    function formatDateTime(val) {
        if (!val && val !== 0) return '';
        var ts = val;
        if (typeof ts === 'string' && /^\d+$/.test(ts)) ts = Number(ts);
        if (typeof ts === 'number' && ts.toString().length === 10) ts = ts * 1000;
        if (typeof ts === 'string') {
            var p = Date.parse(ts);
            if (!isNaN(p)) ts = p;
            else {
                var alt = ts.replace(' ', 'T');
                p = Date.parse(alt);
                if (!isNaN(p)) ts = p;
                else return ts;
            }
        }
        var d = new Date(ts);
        if (isNaN(d.getTime())) return val;
        var dateStr = d.getFullYear() + '-' + pad(d.getMonth()+1) + '-' + pad(d.getDate()) + ' ' + pad(d.getHours()) + ':' + pad(d.getMinutes()) + ':' + pad(d.getSeconds());

        // 时区信息（本地时区）
        var offsetMin = -d.getTimezoneOffset();
        var sign = offsetMin >= 0 ? '+' : '-';
        var absMin = Math.abs(offsetMin);
        var hh = Math.floor(absMin / 60);
        var mm = absMin % 60;

        try {
            var iana = Intl.DateTimeFormat().resolvedOptions().timeZone;
            if (iana) {
                var parts = iana.split('/');
                var region = parts[0] || '';
                var city = parts.slice(1).join('/') || '';
                var regionMap = {
                    'Asia': '亚洲',
                    'Europe': '欧洲',
                    'America': '美洲',
                    'Africa': '非洲',
                    'Atlantic': '大西洋',
                    'Pacific': '太平洋',
                    'Indian': '印度洋',
                    'Arctic': '北极',
                    'Antarctica': '南极'
                };
                var regionZh = regionMap[region] || region;
                var cityKey = (city || '').replace(/_/g, ' ').toLowerCase();
                var cityMap = {
                    'shanghai': '上海',
                    'beijing': '北京',
                    'tokyo': '东京',
                    'new york': '纽约',
                    'los angeles': '洛杉矶',
                    'kolkata': '加尔各答',
                    'mumbai': '孟买',
                    'bangkok': '曼谷',
                    'berlin': '柏林',
                    'kyiv': '基辅',
                    'kiev': '基辅',
                    'singapore': '新加坡',
                    'london': '伦敦',
                    'paris': '巴黎',
                    'sydney': '悉尼',
                    'auckland': '奥克兰'
                };
                var cityZh = cityMap[cityKey] || city;
                var display = (regionZh ? regionZh + '/' : '') + cityZh;
                return display && display !== '/' ? (dateStr + ' (' + display + ')') : (dateStr + ' (' + iana + ')');
            }
        } catch (e) {
            // ignore
        }

        var offsetKey = (sign === '+' ? '+' : '-') + (hh < 10 ? '0' + hh : '' + hh) + (mm === 30 ? '.5' : '');
        var offsetToIana = {
            '+08': 'Asia/Shanghai',
            '+00': 'Etc/UTC',
            '-05': 'America/New_York',
            '-08': 'America/Los_Angeles',
            '+09': 'Asia/Tokyo',
            '+05.5': 'Asia/Kolkata',
            '+07': 'Asia/Bangkok',
            '+01': 'Europe/Berlin',
            '+02': 'Europe/Kiev'
        };
        var ianaFallback = offsetToIana[offsetKey];
        if (ianaFallback) {
            var parts2 = ianaFallback.split('/');
            var region2 = parts2[0] || '';
            var city2 = parts2.slice(1).join('/') || '';
            var regionMap2 = { 'Asia': '亚洲', 'Europe': '欧洲', 'America': '美洲' };
            var regionZh2 = regionMap2[region2] || region2;
            return dateStr + ' (' + (regionZh2 ? regionZh2 + '/' + city2 : ianaFallback) + ')';
        }

        return dateStr;
    }

    // 导出到全局，页面可以直接调用 formatDateTime(val)
    window.formatDateTime = formatDateTime;
})();
