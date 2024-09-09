// SEO Analyzer Tag
(function() {
    var sessionId = Math.random().toString(36).substring(2, 15) + Math.random().toString(36).substring(2, 15);
    var startTime = new Date().getTime();
    var userId = window.seoAnalyzerUserId;
    var projectId = window.seoAnalyzerProjectId;

    if (!userId || !projectId) {
        console.error('SEO Analyzer: userId and projectId must be set');
        return;
    }

    window.addEventListener('beforeunload', function() {
        var endTime = new Date().getTime();
        var duration = endTime - startTime;

        var data = {
            sessionId: sessionId,
            userId: userId,
            projectId: projectId,
            url: window.location.href,
            duration: duration,
            screenWidth: window.screen.width,
            screenHeight: window.screen.height,
            userAgent: navigator.userAgent
        };

        navigator.sendBeacon('/api/collect-session', JSON.stringify(data));
    });
})();
