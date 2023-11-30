yum install git

move the private key file `rd_rsa` to the directory `~/.ssh/`  
`ssh -T git@github.com` // test connection
`git config --global user.name "your name"` // set user name  
`git config --global user.email "your email"` // set user email  
`git config --list` // check config  
`git config --global url."git@github.com:".insteadOf "https://github.com/"` // use ssh instead of https
`git config --global --unset url."git@github.com:".insteadOf` // unset ssh instead of https