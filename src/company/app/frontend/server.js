// Authors:
// 	- Jared Paubel
//  - RooCode agent - local qwen/qwen3.5-9b
// Percentage written by Agent: 100%

const http = require('http');
const fs = require('fs');
const path = require('path');

const FRONTEND_PORT = 3000;
const BACKEND_URL = 'http://localhost:8080';

const server = http.createServer((req, res) => {
    // API proxy routes - forward to backend server
    const apiRoutes = ['/companies', '/departments', '/companies/', '/departments/'];
    
    if (apiRoutes.some(route => req.url.startsWith(route))) {
        // Proxy API request to backend
        const parsedUrl = new URL(req.url, `http://${req.headers.host}`);
        const backendPath = parsedUrl.pathname + parsedUrl.search;
        
        const options = {
            hostname: 'localhost',
            port: 8080,
            path: backendPath,
            method: req.method,
            headers: req.headers
        };
        
        const backendReq = http.request(options, (backendRes) => {
            res.writeHead(backendRes.statusCode, backendRes.headers);
            backendRes.on('data', (chunk) => {
                res.write(chunk);
            });
            backendRes.on('end', () => {
                res.end();
            });
        });
        
        backendReq.on('error', (err) => {
            res.writeHead(502, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify({ error: 'Backend server unavailable', details: err.message }));
        });
        
        req.on('data', (chunk) => {
            backendReq.write(chunk);
        });
        
        req.on('end', () => {
            backendReq.end();
        });
        
        return;
    }
    
    // Static file serving
    let filePath = path.join(__dirname, 'index.html');
    if (req.url !== '/') {
        filePath = path.join(__dirname, req.url);
    }
    
    const ext = path.extname(filePath).toLowerCase();
    const mimeTypes = {
        '.html': 'text/html',
        '.js': 'text/javascript',
        '.css': 'text/css',
        '.json': 'application/json',
        '.png': 'image/png',
        '.jpg': 'image/jpeg',
        '.gif': 'image/gif',
        '.svg': 'image/svg+xml'
    };
    
    const contentType = mimeTypes[ext] || 'text/plain';
    
    fs.readFile(filePath, (err, content) => {
        if (err) {
            res.writeHead(404);
            res.end('File not found');
        } else {
            res.writeHead(200, { 'Content-Type': contentType });
            res.end(content);
        }
    });
});

server.listen(FRONTEND_PORT, () => {
    process.stdout.write(`Frontend server running on http://localhost:${FRONTEND_PORT}\n`);
    process.stdout.write(`API requests proxied to ${BACKEND_URL}\n`);
});
