package swaggercss

import "html/template"

var DefaultCSS template.CSS = `
	body {
		margin: 0;
	} 
	body > #swagger-ui {
		min-height: 100vh;
		background-color: #ebebeb;
		padding-bottom: 30px;
	}
	.swagger-ui .topbar .download-url-wrapper input[type=text] {
		border-color: #FF3E00; 
	}
	.swagger-ui .topbar .download-url-wrapper .download-url-button {
		background: #FF3E00;
	}
	.swagger-ui .models {
		margin-bottom: 0;
	}

	.topbar-wrapper img {
		content: url('/api/asset/logo-branding')
	}
`
