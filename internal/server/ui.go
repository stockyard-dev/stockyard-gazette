package server

var dashboardHTML = []byte(`<!DOCTYPE html>
<html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1.0"><title>Stockyard Gazette</title><style>:root{--bg:#1a1410;--surface:#241c15;--border:#3d2e1e;--rust:#c4622d;--cream:#f5e6c8;--muted:#7a6550;--text:#e8d5b0}*{box-sizing:border-box;margin:0;padding:0}body{background:var(--bg);color:var(--text);font-family:'JetBrains Mono',monospace,sans-serif}header{background:var(--surface);border-bottom:1px solid var(--border);padding:1rem 2rem;display:flex;align-items:center;gap:1rem}.logo{color:var(--rust);font-size:1.25rem;font-weight:700}.badge{background:var(--rust);color:var(--cream);font-size:0.65rem;padding:0.2rem 0.5rem;border-radius:3px;font-weight:600;text-transform:uppercase}main{max-width:1100px;margin:0 auto;padding:2rem}.stats{display:grid;grid-template-columns:repeat(3,1fr);gap:1rem;margin-bottom:2rem}.stat{background:var(--surface);border:1px solid var(--border);border-radius:6px;padding:1.25rem;text-align:center}.stat-value{font-size:1.75rem;font-weight:700;color:var(--rust)}.stat-label{font-size:0.75rem;color:var(--muted);margin-top:0.25rem;text-transform:uppercase;letter-spacing:0.05em}.layout{display:grid;grid-template-columns:1fr 360px;gap:1.5rem}.card{background:var(--surface);border:1px solid var(--border);border-radius:6px;padding:1.5rem;margin-bottom:1rem}.card h2{font-size:0.85rem;color:var(--muted);text-transform:uppercase;letter-spacing:0.08em;margin-bottom:1rem}.form-row{display:flex;gap:0.5rem;margin-bottom:0.75rem;flex-wrap:wrap}select,input,textarea{background:var(--bg);border:1px solid var(--border);color:var(--text);padding:0.5rem 0.75rem;border-radius:4px;font-family:inherit;font-size:0.85rem;flex:1}textarea{resize:vertical;flex:none;width:100%}.btn{background:var(--rust);color:var(--cream);border:none;padding:0.5rem 1rem;border-radius:4px;cursor:pointer;font-family:inherit;font-size:0.85rem;font-weight:600}.btn:hover{opacity:0.85}.btn-sm{padding:0.25rem 0.6rem;font-size:0.75rem}.btn-danger{background:#7a2020}.btn-outline{background:transparent;border:1px solid var(--rust);color:var(--rust)}.post-item{padding:0.75rem;border-bottom:1px solid var(--border);cursor:pointer}.post-item:hover{background:rgba(196,98,45,0.08)}.post-title{font-weight:600;color:var(--cream);font-size:0.9rem}.post-meta{font-size:0.72rem;color:var(--muted);margin-top:0.2rem}.badge-draft{background:#2a2a1a;color:#b8a060;border:1px solid #4a4a2a;border-radius:3px;padding:0.1rem 0.4rem;font-size:0.7rem}.badge-pub{background:#1a3a1a;color:#5cb85c;border:1px solid #2d5a2d;border-radius:3px;padding:0.1rem 0.4rem;font-size:0.7rem}.empty{color:var(--muted);font-size:0.85rem;padding:1rem 0;text-align:center}.pub-body{background:var(--bg);border:1px solid var(--border);border-radius:4px;padding:1rem;font-size:0.85rem;line-height:1.6;white-space:pre-wrap;max-height:300px;overflow-y:auto}</style></head>
<body>
<header><span class="logo">&#x2B21; Stockyard</span><span style="color:var(--muted)">/</span><span style="color:var(--cream);font-weight:600">Gazette</span><span class="badge">Blog</span></header>
<main>
<div class="stats"><div class="stat"><div class="stat-value" id="s1">0</div><div class="stat-label">Posts</div></div><div class="stat"><div class="stat-value" id="s2">0</div><div class="stat-label">Total Views</div></div><div class="stat"><div class="stat-value" id="s3">FREE</div><div class="stat-label">Tier</div></div></div>
<div class="layout">
<div>
<div class="card">
<div style="display:flex;justify-content:space-between;align-items:center;margin-bottom:1rem"><h2 style="margin:0" id="editor-title">New Post</h2><button class="btn-sm btn-outline btn" onclick="resetEditor()">+ New</button></div>
<div class="form-row"><input id="f-title" placeholder="Post title"><select id="f-status"><option value="draft">Draft</option><option value="published">Published</option></select></div>
<div class="form-row"><input id="f-tags" placeholder="Tags (comma separated)"><input id="f-excerpt" placeholder="Excerpt (optional)"></div>
<textarea id="f-body" rows="12" placeholder="Write your post here..."></textarea>
<div style="display:flex;gap:0.5rem;margin-top:0.75rem"><button class="btn" onclick="savePost()">Save Post</button><button class="btn btn-outline" id="pub-btn" onclick="publishPost()" style="display:none">Publish</button></div>
</div>
</div>
<div>
<div class="card"><h2>Posts</h2>
<div class="form-row"><input id="f-search" placeholder="Search..." oninput="searchPosts()"><select id="f-filter" onchange="loadPosts()"><option value="">All</option><option value="draft">Drafts</option><option value="published">Published</option></select></div>
<div id="post-list"><div class="empty">No posts yet</div></div>
</div>
</div>
</div>
</main>
<script>
var editId=null;
function load(){fetch('/api/stats').then(function(r){return r.json()}).then(function(d){document.getElementById('s1').textContent=d.posts||0;document.getElementById('s2').textContent=d.total_views||0})}
function loadPosts(){var s=document.getElementById('f-filter').value;fetch('/api/posts'+(s?'?status='+s:'')).then(function(r){return r.json()}).then(renderPosts)}
function renderPosts(list){var el=document.getElementById('post-list');el.innerHTML=list.length?list.map(function(p){return'<div class="post-item" onclick="editPost('+JSON.stringify(p).replace(/'/g,"\\x27")+')"><div style="display:flex;justify-content:space-between"><span class="post-title">'+p.title+'</span><span class="badge-'+(p.status==='published'?'pub':'draft'))
+'">'+p.status+'</span></div><div class="post-meta">&#x1F441; '+p.views+' &bull; '+p.created_at.substring(0,10)+'</div></div>'}).join(''):'<div class="empty">No posts</div>'}
function searchPosts(){var q=document.getElementById('f-search').value.trim();if(!q){loadPosts();return};fetch('/api/search?q='+encodeURIComponent(q)).then(function(r){return r.json()}).then(renderPosts)}
function editPost(p){editId=p.id;document.getElementById('editor-title').textContent='Edit Post';document.getElementById('f-title').value=p.title;document.getElementById('f-status').value=p.status;document.getElementById('f-tags').value=p.tags||"";document.getElementById('f-excerpt').value=p.excerpt||"";document.getElementById('f-body').value=p.body||"";document.getElementById('pub-btn').style.display=p.status==='draft'?'inline-block':'none'}
function resetEditor(){editId=null;document.getElementById('editor-title').textContent='New Post';['f-title','f-tags','f-excerpt','f-body'].forEach(function(id){document.getElementById(id).value=''});document.getElementById('f-status').value='draft';document.getElementById('pub-btn').style.display='none'}
function savePost(){var d={title:document.getElementById('f-title').value.trim(),status:document.getElementById('f-status').value,tags:document.getElementById('f-tags').value.trim(),excerpt:document.getElementById('f-excerpt').value.trim(),body:document.getElementById('f-body').value.trim()};if(!d.title)return;var url=editId?'/api/posts/'+editId:'/api/posts';var method=editId?'PUT':'POST';fetch(url,{method:method,headers:{'Content-Type':'application/json'},body:JSON.stringify(d)}).then(function(r){return r.json()}).then(function(p){if(!editId)editId=p.id;loadPosts();load()})}
function publishPost(){document.getElementById('f-status').value='published';savePost();document.getElementById('pub-btn').style.display='none'}
load();loadPosts();
</script></body></html>`)
