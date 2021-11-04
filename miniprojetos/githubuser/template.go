package main

import "text/template"

var userTemplate = template.Must(template.New("user").Parse(`
<h2>{{.Name}}, {{.Login}}</h2>
<img src='{{.AvatarURL}}'>
<h3>User info</h3>
<ul style="list-style-type: none;">
	<li>ID {{.ID}}</li>
	<li>NodeID {{.NodeID}}</li>
	<li>GravatarID {{.GravatarID}}</li>
	<li>Type {{.Type}}</li>
	<li>SiteAdmin {{.SiteAdmin}}</li>
	<li>Company {{.Company}}</li>
	<li>Blog {{.Blog}}</li>
	<li>Location {{.Location}}</li>
	<li>Email {{.Email}}</li>
	<li>Hireable {{.Hireable}}</li>
	<li>Bio {{.Bio}}</li>
	<li>TwitterUsername {{.TwitterUsername}}</li>
	<li>PublicRepos {{.PublicRepos}}</li>
	<li>PublicGists {{.PublicGists}}</li>
	<li>Followers {{.Followers}}</li>
	<li>Following {{.Following}}</li>
	<li>CreatedAt {{.CreatedAt}}</li>
	<li>UpdatedAt {{.UpdatedAt}}</li>
</ul>
`))
