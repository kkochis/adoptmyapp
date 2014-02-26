package robot

import (
	"testing"
)

func expectString(t *testing.T, a, b string) {
	if a != b {
		t.Errorf("Unexpected string: %s != %s", a, b)
	}
}

func expectInt64(t *testing.T, a, b int64) {
	if a != b {
		t.Errorf("Unexpected integer: %d != %d", a, b)
	}
}

func TestParseRepository(t *testing.T) {
	// Test the parsing of a typical github repository API request
	repo, err := ParseRepository([]byte(repository))
	if err != nil {
		t.Fatal(err)
	}
	// Repo details
	expectString(t, repo.Name, "adoptmyapp")
	expectInt64(t, repo.Id, 17196911)

	// Owner details
	expectString(t, repo.Owner.Login, "kkochis")
	expectInt64(t, repo.Owner.Id, 648320)
}

func TestConvertRepoURL(t *testing.T) {
	apiURL, err := ConvertRepoURL(`https://github.com/kkochis/adoptmyapp`)
	if err != nil {
		t.Fatal(err)
	}
	expectString(t, apiURL, `https://api.github.com/repos/kkochis/adoptmyapp`)
}

var repository = `{
  "id": 17196911,
  "name": "adoptmyapp",
  "full_name": "kkochis/adoptmyapp",
  "owner": {
    "login": "kkochis",
    "id": 648320,
    "avatar_url": "https://gravatar.com/avatar/395a7672355d261c60bab95d44383f84?d=https%3A%2F%2Fidenticons.github.com%2F5c9ffa7ceff30daffa97461e2cfcdac7.png&r=x",
    "gravatar_id": "395a7672355d261c60bab95d44383f84",
    "url": "https://api.github.com/users/kkochis",
    "html_url": "https://github.com/kkochis",
    "followers_url": "https://api.github.com/users/kkochis/followers",
    "following_url": "https://api.github.com/users/kkochis/following{/other_user}",
    "gists_url": "https://api.github.com/users/kkochis/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/kkochis/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/kkochis/subscriptions",
    "organizations_url": "https://api.github.com/users/kkochis/orgs",
    "repos_url": "https://api.github.com/users/kkochis/repos",
    "events_url": "https://api.github.com/users/kkochis/events{/privacy}",
    "received_events_url": "https://api.github.com/users/kkochis/received_events",
    "type": "User",
    "site_admin": false
  },
  "private": false,
  "html_url": "https://github.com/kkochis/adoptmyapp",
  "description": "Adopt My App is an app that helps people share and rescue abandoned apps and hackathon projects",
  "fork": false,
  "url": "https://api.github.com/repos/kkochis/adoptmyapp",
  "forks_url": "https://api.github.com/repos/kkochis/adoptmyapp/forks",
  "keys_url": "https://api.github.com/repos/kkochis/adoptmyapp/keys{/key_id}",
  "collaborators_url": "https://api.github.com/repos/kkochis/adoptmyapp/collaborators{/collaborator}",
  "teams_url": "https://api.github.com/repos/kkochis/adoptmyapp/teams",
  "hooks_url": "https://api.github.com/repos/kkochis/adoptmyapp/hooks",
  "issue_events_url": "https://api.github.com/repos/kkochis/adoptmyapp/issues/events{/number}",
  "events_url": "https://api.github.com/repos/kkochis/adoptmyapp/events",
  "assignees_url": "https://api.github.com/repos/kkochis/adoptmyapp/assignees{/user}",
  "branches_url": "https://api.github.com/repos/kkochis/adoptmyapp/branches{/branch}",
  "tags_url": "https://api.github.com/repos/kkochis/adoptmyapp/tags",
  "blobs_url": "https://api.github.com/repos/kkochis/adoptmyapp/git/blobs{/sha}",
  "git_tags_url": "https://api.github.com/repos/kkochis/adoptmyapp/git/tags{/sha}",
  "git_refs_url": "https://api.github.com/repos/kkochis/adoptmyapp/git/refs{/sha}",
  "trees_url": "https://api.github.com/repos/kkochis/adoptmyapp/git/trees{/sha}",
  "statuses_url": "https://api.github.com/repos/kkochis/adoptmyapp/statuses/{sha}",
  "languages_url": "https://api.github.com/repos/kkochis/adoptmyapp/languages",
  "stargazers_url": "https://api.github.com/repos/kkochis/adoptmyapp/stargazers",
  "contributors_url": "https://api.github.com/repos/kkochis/adoptmyapp/contributors",
  "subscribers_url": "https://api.github.com/repos/kkochis/adoptmyapp/subscribers",
  "subscription_url": "https://api.github.com/repos/kkochis/adoptmyapp/subscription",
  "commits_url": "https://api.github.com/repos/kkochis/adoptmyapp/commits{/sha}",
  "git_commits_url": "https://api.github.com/repos/kkochis/adoptmyapp/git/commits{/sha}",
  "comments_url": "https://api.github.com/repos/kkochis/adoptmyapp/comments{/number}",
  "issue_comment_url": "https://api.github.com/repos/kkochis/adoptmyapp/issues/comments/{number}",
  "contents_url": "https://api.github.com/repos/kkochis/adoptmyapp/contents/{+path}",
  "compare_url": "https://api.github.com/repos/kkochis/adoptmyapp/compare/{base}...{head}",
  "merges_url": "https://api.github.com/repos/kkochis/adoptmyapp/merges",
  "archive_url": "https://api.github.com/repos/kkochis/adoptmyapp/{archive_format}{/ref}",
  "downloads_url": "https://api.github.com/repos/kkochis/adoptmyapp/downloads",
  "issues_url": "https://api.github.com/repos/kkochis/adoptmyapp/issues{/number}",
  "pulls_url": "https://api.github.com/repos/kkochis/adoptmyapp/pulls{/number}",
  "milestones_url": "https://api.github.com/repos/kkochis/adoptmyapp/milestones{/number}",
  "notifications_url": "https://api.github.com/repos/kkochis/adoptmyapp/notifications{?since,all,participating}",
  "labels_url": "https://api.github.com/repos/kkochis/adoptmyapp/labels{/name}",
  "releases_url": "https://api.github.com/repos/kkochis/adoptmyapp/releases{/id}",
  "created_at": "2014-02-26T03:00:56Z",
  "updated_at": "2014-02-26T03:00:56Z",
  "pushed_at": "2014-02-26T03:00:56Z",
  "git_url": "git://github.com/kkochis/adoptmyapp.git",
  "ssh_url": "git@github.com:kkochis/adoptmyapp.git",
  "clone_url": "https://github.com/kkochis/adoptmyapp.git",
  "svn_url": "https://github.com/kkochis/adoptmyapp",
  "homepage": null,
  "size": 0,
  "stargazers_count": 1,
  "watchers_count": 1,
  "language": null,
  "has_issues": true,
  "has_downloads": true,
  "has_wiki": true,
  "forks_count": 0,
  "mirror_url": null,
  "open_issues_count": 0,
  "forks": 0,
  "open_issues": 0,
  "watchers": 1,
  "default_branch": "master",
  "master_branch": "master",
  "network_count": 0,
  "subscribers_count": 2
}`