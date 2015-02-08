var pathFn = require('path');
var util = require('hexo-util');
var fs = require('hexo-fs');
var spawn = util.spawn;

var passYamlDir = pathFn.join(__dirname, 'templates/pass');
var pubYamlDir = pathFn.join(__dirname, 'templates/pub');

var publicDir = hexo.config.public_dir || './public';
var baseDir = hexo.base_dir;

var deployDir = pathFn.join(baseDir, '.deploy_appengine');
var deployPubDir = pathFn.join(deployDir, 'static');

var log = hexo.log;

hexo.extend.deployer.register('appengine', function(args) {

  var project = args.project;
  var version = args.version || "master";
  var verbose = !args.silent;
  var password = args.password;
  var deploy = !args.dryrun;

  if (!args.project) {
    var help = '';

    help += 'You have to configure the deployment settings in _config.yml first!\n\n';
    help += 'Example:\n';
    help += '  deploy:\n';
    help += '    type: appengine\n';
    help += '    project: <project id>\n';
    help += '    version: <versionname defaults to master>\n';
    help += '    password: <optional to make password protected>\n\n';
    help += 'For more help, you can check the docs: http://hexo.io/docs/deployment.html';

    console.log(help);
    return;
  }

  return fs.exists(deployDir).then(function(exists) {
    if (!exists) return;
    log.info('Clearing old .deploy_appengine folder...');
    return fs.emptyDir(deployDir);
  }).then(function() {
    log.info('Setting up AppEngine deployment...');
    if (password) {
      return fs.copyDir(passYamlDir, deployDir).then(function() {
        return fs.writeFile(pathFn.join(deployDir, 'password'), password);
      });
    }
    return fs.copyDir(pubYamlDir, deployDir);
  }).then(function() {
    log.info('Copying files from public folder...');
    return fs.copyDir(publicDir, deployPubDir);
  }).then(function() {
    if (deploy) {
      return spawn('appcfg.py', ['update', '-A', project, '-V', version, '.'], {
        cwd: deployDir,
        verbose: verbose
      }).then(function() {
        log.info('Deployed to https://' + version + '-dot-' + project + '.appspot.com');
      });
    } else {
      log.info('Dry run: files written to ' + deployDir +
      '. Can check local dev with dev_appserver.py ' + deployDir);
    }
  });

});
