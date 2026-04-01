package server
import("encoding/json";"net/http";"strconv";"github.com/stockyard-dev/stockyard-gazette/internal/store")
func(s *Server)handleListPosts(w http.ResponseWriter,r *http.Request){status:=r.URL.Query().Get("status");list,_:=s.db.ListPosts(status);if list==nil{list=[]store.Post{}};writeJSON(w,200,list)}
func(s *Server)handleGetPost(w http.ResponseWriter,r *http.Request){slug:=r.PathValue("slug");p,err:=s.db.GetPost(slug);if err!=nil{writeError(w,404,"not found");return};writeJSON(w,200,p)}
func(s *Server)handleCreatePost(w http.ResponseWriter,r *http.Request){var p store.Post;json.NewDecoder(r.Body).Decode(&p);if p.Title==""{writeError(w,400,"title required");return};if err:=s.db.CreatePost(&p);err!=nil{writeError(w,500,err.Error());return};writeJSON(w,201,p)}
func(s *Server)handleUpdatePost(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);var p store.Post;json.NewDecoder(r.Body).Decode(&p);p.ID=id;if p.Title==""{writeError(w,400,"title required");return};if err:=s.db.UpdatePost(&p);err!=nil{writeError(w,500,err.Error());return};writeJSON(w,200,p)}
func(s *Server)handleDeletePost(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.DeletePost(id);writeJSON(w,200,map[string]string{"status":"deleted"})}
func(s *Server)handleSearch(w http.ResponseWriter,r *http.Request){q:=r.URL.Query().Get("q");if q==""{writeError(w,400,"q required");return};list,_:=s.db.SearchPosts(q);if list==nil{list=[]store.Post{}};writeJSON(w,200,list)}
func(s *Server)handleStats(w http.ResponseWriter,r *http.Request){n,_:=s.db.CountPosts();v,_:=s.db.TotalViews();writeJSON(w,200,map[string]interface{}{"posts":n,"total_views":v})}
