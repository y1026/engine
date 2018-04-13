package service

import (
	"fmt"
	"os/user"
	"strings"

	"os"

	"github.com/it-chain/it-chain-Engine/itcode/domain/itcode"
	"github.com/pkg/errors"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

var tmp = "./.tmp"

type GitApi struct {
	sshAuth             *ssh.PublicKeys
	defaultBackUpGitUrl string
}

//todo get defaultBackUpGitUrl from config
func NewGitApi() GitApi {

	currentUser, err := user.Current()

	if err != nil {
		panic(fmt.Sprintf("fail to init GitApi [%s]", err.Error()))
	}

	sshAuth, err := ssh.NewPublicKeysFromFile("git", currentUser.HomeDir+"/.ssh/id_rsa", "")

	if err != nil {
		panic(fmt.Sprintf("fail to init GitApi [%s]", err.Error()))
	}

	return GitApi{
		sshAuth:             sshAuth,
		defaultBackUpGitUrl: "https://github.com/steve-buzzni/",
	}
}

//get itcode from outside
//todo SSH ENV로 ssh key 불러오기
func (g GitApi) Clone(gitUrl string) (*itcode.ItCode, error) {

	name := getNameFromGitUrl(gitUrl)

	if name == "" {
		return nil, errors.New(fmt.Sprintf("Invalid url name [%s]", gitUrl))
	}

	r, err := git.PlainClone(tmp+"/"+name, false, &git.CloneOptions{
		URL:               gitUrl,
		Auth:              g.sshAuth,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	if err != nil {
		return nil, err
	}

	head, err := r.Head()

	if err != nil {
		return nil, err
	}

	lastHeadCommit, err := r.CommitObject(head.Hash())
	commitHash := lastHeadCommit.Hash.String()

	if err != nil {
		return nil, err
	}

	//todo os separator
	sc := itcode.NewItCode(name, gitUrl, tmp+"/"+name, commitHash)

	return sc, nil
}

//change remote origin and push code to auth/backup repo
func (g GitApi) Push(itCode *itcode.ItCode) error {
	itCodePath := itCode.Path

	if !dirExists(itCodePath) {
		return errors.New(fmt.Sprintf("Invalid itCode Path [%s]", itCodePath))
	}

	err := g.ChangeRemote(itCodePath, g.defaultBackUpGitUrl+itCode.RepositoryName)

	if err != nil {
		return err
	}

	err = g.push(itCodePath)

	if err != nil {
		return err
	}

	return nil
}

//change origin remote
func (g GitApi) ChangeRemote(path string, originUrl string) error {

	r, err := git.PlainOpen(path)

	if err != nil {
		return err
	}

	err = r.DeleteRemote(git.DefaultRemoteName)

	if err != nil {
		return err
	}

	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: git.DefaultRemoteName,
		URLs: []string{originUrl},
	})

	if err != nil {
		return err
	}

	return nil
}

//push to backup server
//todo get username and password from config
func (g GitApi) push(path string) error {

	au := &http.BasicAuth{Username: "steve@buzzni.com", Password: "itchain123"}

	r, err := git.PlainOpen(path)

	if err != nil {
		return err
	}

	err = r.Push(&git.PushOptions{
		RemoteName: git.DefaultRemoteName,
		Auth:       au,
	})

	if err != nil {
		return err
	}

	return nil
}

//get repo name from git url
func getNameFromGitUrl(gitUrl string) string {
	parsed := strings.Split(gitUrl, "/")

	if len(parsed) == 0 {
		return ""
	}

	name := strings.Split(parsed[len(parsed)-1], ".")

	if len(name) == 0 {
		return ""
	}

	return name[0]
}

//check whetehr dir is exist or not
func dirExists(path string) bool {
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return true
	}
	return false
}
