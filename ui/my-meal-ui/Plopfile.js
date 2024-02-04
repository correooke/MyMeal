module.exports = function (plop) {
    // Component generator
    plop.setGenerator('component', {
        description: 'Create a React component',
        prompts: [{
            type: 'input',
            name: 'name',
            message: 'Component name',
            default: 'MyComponent',
        },
        {
            type: 'input',
            name: 'folder',
            message: 'Folder path',
            default: 'MyComponent',
        }],
        actions: [{
            type: 'add',
            path: 'components/{{folder}}/{{name}}.tsx',
            templateFile: 'plop-templates/component.hbs',
        }]
    });

    plop.setGenerator('page', {
        description: 'Create a Next.js page',
        prompts: [{
            type: 'input',
            name: 'name',
            message: 'Page name',
            default: 'MyPage',
        }],
        actions: [{
            type: 'add',
            path: 'pages/{{kebabCase name}}.js',
            templateFile: 'plop-templates/page.hbs',
        }]
    });
    
};
