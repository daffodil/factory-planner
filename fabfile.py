# -*- coding: utf-8 -*-
# author: pete@daffodil.uk.com

import os
import sys

#import configparser

from fabric.api import env, local, run, cd, sudo, prompt
from fabric import colors


env.hosts = ['factory-planner.daffodil.uk.com']
env.user = "facman"
env.password = "using-ssh-ssl-key"
env.use_ssh_config = True # this is using ~/.ssh/config = sshkey login

TEST_ROOT = "/home/facman/factory-planner"

def d_test():
	"""Deploys and restarts test server"""
	local("git push origin --all")
	with cd(TEST_ROOT):
		run("git pull")