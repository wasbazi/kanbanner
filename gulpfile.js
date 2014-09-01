var gulp = require('gulp')
var uglify = require('gulp-uglify')
var concat = require('gulp-concat')

var scripts = './public/javascript/angular/**/*.js'
var dest = './public/javascript'
var output = 'compiled.js'

gulp.task('javascript', function(){
  gulp.src(scripts)
    .pipe(uglify())
    .pipe(concat(output))
    .pipe(gulp.dest(dest))
})

gulp.task('javascript-dev', function(){
  gulp.src(scripts)
    .pipe(concat(output))
    .pipe(gulp.dest(dest))
})

gulp.task('watch', function(){
  gulp.watch('./public/**/*.js', ['javascript-dev'])
})
