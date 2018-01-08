'use strict';

const Generator = require('yeoman-generator');

module.exports = class extends Generator {
  prompting() {
    const prompts = [{
      type: 'input',
      name: 'package',
      message: 'package path (ex. github.com/suzuki-shunsuke/git-rm-branch)'
    }];

    return this.prompt(prompts).then(answers => {
      this.answers = answers;
    });
  }

  writing() {
    ['main.go', 'cmds/init.go', 'services/util.go', 'Makefile'].forEach(src => {
      this.fs.copyTpl(
        this.templatePath(src), this.destinationPath(src),
        {package: this.answers.package});
    });
    this.fs.copy(
      this.templatePath('gitignore'),
      this.destinationPath('.gitignore'));
    ['data'].forEach(key => {
      this.fs.copy(
        this.templatePath(key),
        this.destinationPath(key));
    });
  }
};
