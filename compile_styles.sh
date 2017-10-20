# !/bin/bash

MARKUP_PATH=$1
PROJECT_ROOT=$(pwd)
cd $MARKUP_PATH
git checkout go-templates
npm install
npm run build
cd $PROJECT_ROOT

rm -rf templates
mkdir templates

rm -rf static
mkdir static
mkdir static/styles
mkdir static/images
mkdir static/fonts

# Updating html templates.
cp "${MARKUP_PATH}/dist/main-page/main-page.html" "${PROJECT_ROOT}/templates/index.html"
cp "${MARKUP_PATH}/dist/login-page/login-page.html" "${PROJECT_ROOT}/templates/login.html"
cp "${MARKUP_PATH}/dist/article-page/article-page.html" "${PROJECT_ROOT}/templates/article.html"

# Images and fonts.
cp -r "${MARKUP_PATH}/dist/images" "${PROJECT_ROOT}/static"
cp -r "${MARKUP_PATH}/dist/fonts" "${PROJECT_ROOT}/static"

# Updating styles.
cp "${MARKUP_PATH}/dist/styles/assets.min.css" "${PROJECT_ROOT}/static/styles/assets.min.css"
cp "${MARKUP_PATH}/dist/styles/bundle.min.css" "${PROJECT_ROOT}/static/styles/bundle.min.css"
