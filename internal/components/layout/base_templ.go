// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.696
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Base() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html lang=\"en\"><head><base href=\"./\"><meta charset=\"utf-8\"><meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0, shrink-to-fit=no\"><meta name=\"description\" content=\"CoreUI - Open Source Bootstrap Admin Template\"><meta name=\"author\" content=\"Łukasz Holeczek\"><meta name=\"keyword\" content=\"Bootstrap,Admin,Template,Open,Source,jQuery,CSS,HTML,RWD,Dashboard\"><title>Portanic Developer Platform</title><link rel=\"apple-touch-icon\" sizes=\"57x57\" href=\"assets/favicon/apple-icon-57x57.png\"><link rel=\"apple-touch-icon\" sizes=\"60x60\" href=\"assets/favicon/apple-icon-60x60.png\"><link rel=\"apple-touch-icon\" sizes=\"72x72\" href=\"assets/favicon/apple-icon-72x72.png\"><link rel=\"apple-touch-icon\" sizes=\"76x76\" href=\"assets/favicon/apple-icon-76x76.png\"><link rel=\"apple-touch-icon\" sizes=\"114x114\" href=\"assets/favicon/apple-icon-114x114.png\"><link rel=\"apple-touch-icon\" sizes=\"120x120\" href=\"assets/favicon/apple-icon-120x120.png\"><link rel=\"apple-touch-icon\" sizes=\"144x144\" href=\"assets/favicon/apple-icon-144x144.png\"><link rel=\"apple-touch-icon\" sizes=\"152x152\" href=\"assets/favicon/apple-icon-152x152.png\"><link rel=\"apple-touch-icon\" sizes=\"180x180\" href=\"assets/favicon/apple-icon-180x180.png\"><link rel=\"icon\" type=\"image/png\" sizes=\"192x192\" href=\"assets/favicon/android-icon-192x192.png\"><link rel=\"icon\" type=\"image/png\" sizes=\"32x32\" href=\"assets/favicon/favicon-32x32.png\"><link rel=\"icon\" type=\"image/png\" sizes=\"96x96\" href=\"assets/favicon/favicon-96x96.png\"><link rel=\"icon\" type=\"image/png\" sizes=\"16x16\" href=\"assets/favicon/favicon-16x16.png\"><link rel=\"manifest\" href=\"assets/favicon/manifest.json\"><meta name=\"msapplication-TileColor\" content=\"#ffffff\"><meta name=\"msapplication-TileImage\" content=\"assets/favicon/ms-icon-144x144.png\"><meta name=\"theme-color\" content=\"#ffffff\"><!-- Vendors styles--><link rel=\"stylesheet\" href=\"/vendors/simplebar/css/simplebar.css\"><link rel=\"stylesheet\" href=\"/css/vendors/simplebar.css\"><!-- Main styles for this application--><link href=\"/css/style.css\" rel=\"stylesheet\"></head><body><div class=\"sidebar sidebar-dark sidebar-fixed border-end\" id=\"sidebar\"><div class=\"sidebar-header border-bottom\"><div class=\"sidebar-brand\">Portanic</div></div><ul class=\"sidebar-nav\" data-coreui=\"navigation\" data-simplebar=\"\"><li class=\"nav-item\"><a class=\"nav-link\" href=\"index.html\">Dashboard</a></li><li class=\"nav-title\">Components</li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/templates\">Service Templates</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/catalog\">Service Catalog</a></li><li class=\"nav-item mt-auto\"><a class=\"nav-link\" href=\"https://docs.portanic.com/\" target=\"_blank\">Docs</a></li></ul><div class=\"sidebar-footer border-top d-none d-md-flex\"><button class=\"sidebar-toggler\" type=\"button\" data-coreui-toggle=\"unfoldable\"></button></div></div><div class=\"wrapper d-flex flex-column min-vh-100\"><header class=\"header header-sticky p-0 mb-4\"><div class=\"container-fluid border-bottom px-4\"><button class=\"header-toggler\" type=\"button\" onclick=\"coreui.Sidebar.getInstance(document.querySelector(&#39;#sidebar&#39;)).toggle()\" style=\"margin-inline-start: -14px;\"><svg class=\"icon icon-lg\"><use xlink:href=\"vendors/@coreui/icons/svg/free.svg#cil-menu\"></use></svg></button><ul class=\"header-nav ms-auto\"><li class=\"nav-item\"><a class=\"nav-link\" href=\"#\"><svg class=\"icon icon-lg\"><use xlink:href=\"vendors/@coreui/icons/svg/free.svg#cil-bell\"></use></svg></a></li></ul><ul class=\"header-nav\"><li class=\"nav-item py-1\"><div class=\"vr h-100 mx-2 text-body text-opacity-75\"></div></li><li class=\"nav-item dropdown\"><button class=\"btn btn-link nav-link py-2 px-2 d-flex align-items-center\" type=\"button\" aria-expanded=\"false\" data-coreui-toggle=\"dropdown\"><svg class=\"icon icon-lg theme-icon-active\"><use xlink:href=\"vendors/@coreui/icons/svg/free.svg#cil-contrast\"></use></svg></button><ul class=\"dropdown-menu dropdown-menu-end\" style=\"--cui-dropdown-min-width: 8rem;\"><li><button class=\"dropdown-item d-flex align-items-center\" type=\"button\" data-coreui-theme-value=\"light\"><svg class=\"icon icon-lg me-3\"><use xlink:href=\"vendors/@coreui/icons/svg/free.svg#cil-sun\"></use></svg>Light</button></li><li><button class=\"dropdown-item d-flex align-items-center\" type=\"button\" data-coreui-theme-value=\"dark\"><svg class=\"icon icon-lg me-3\"><use xlink:href=\"vendors/@coreui/icons/svg/free.svg#cil-moon\"></use></svg>Dark</button></li><li><button class=\"dropdown-item d-flex align-items-center active\" type=\"button\" data-coreui-theme-value=\"auto\"><svg class=\"icon icon-lg me-3\"><use xlink:href=\"vendors/@coreui/icons/svg/free.svg#cil-contrast\"></use></svg>Auto</button></li></ul></li><li class=\"nav-item py-1\"><div class=\"vr h-100 mx-2 text-body text-opacity-75\"></div></li><li class=\"nav-item dropdown\"><a class=\"nav-link py-0 pe-0\" data-coreui-toggle=\"dropdown\" href=\"#\" role=\"button\" aria-haspopup=\"true\" aria-expanded=\"false\"><div class=\"avatar avatar-md\"><img class=\"avatar-img\" src=\"assets/img/avatars/8.jpg\" alt=\"user@email.com\"></div></a><div class=\"dropdown-menu dropdown-menu-end pt-0\"><div class=\"dropdown-header bg-body-tertiary text-body-secondary fw-semibold my-2\"><div class=\"fw-semibold\">Settings</div></div><a class=\"dropdown-item\" href=\"#\"><svg class=\"icon me-2\"><use xlink:href=\"vendors/@coreui/icons/svg/free.svg#cil-user\"></use></svg> Profile</a> <a class=\"dropdown-item\" href=\"#\"><svg class=\"icon me-2\"><use xlink:href=\"vendors/@coreui/icons/svg/free.svg#cil-settings\"></use></svg> Settings</a> <a class=\"dropdown-item\" href=\"#\"><svg class=\"icon me-2\"><use xlink:href=\"vendors/@coreui/icons/svg/free.svg#cil-credit-card\"></use></svg> Payments<span class=\"badge badge-sm bg-secondary ms-2\">42</span></a> <a class=\"dropdown-item\" href=\"#\"><svg class=\"icon me-2\"><use xlink:href=\"vendors/@coreui/icons/svg/free.svg#cil-file\"></use></svg> Projects<span class=\"badge badge-sm bg-primary ms-2\">42</span></a><div class=\"dropdown-divider\"></div><a class=\"dropdown-item\" href=\"#\"><svg class=\"icon me-2\"><use xlink:href=\"vendors/@coreui/icons/svg/free.svg#cil-lock-locked\"></use></svg> Lock Account</a> <a class=\"dropdown-item\" href=\"#\"><svg class=\"icon me-2\"><use xlink:href=\"vendors/@coreui/icons/svg/free.svg#cil-account-logout\"></use></svg> Logout</a></div></li></ul></div><div class=\"container-fluid px-4\"><nav aria-label=\"breadcrumb\"><ol class=\"breadcrumb my-0\"><li class=\"breadcrumb-item\"><a href=\"#\">Home</a></li><li class=\"breadcrumb-item active\"><span>Dashboard</span></li></ol></nav></div></header><div class=\"body flex-grow-1\"><div class=\"container-lg px-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><footer class=\"footer px-4\"><div>© 2024 portanic.</div></footer></div><!-- CoreUI and necessary plugins--><script src=\"https://unpkg.com/htmx.org@1.9.12\" integrity=\"sha384-ujb1lZYygJmzgSwoxRggbCHcjc0rB2XoQrxeTUQyRjrOnlCoYta87iKBWq3EsdM2\" crossorigin=\"anonymous\"></script><script src=\"vendors/@coreui/coreui/js/coreui.bundle.min.js\"></script><script src=\"vendors/simplebar/js/simplebar.min.js\"></script><script>\n    const header = document.querySelector('header.header');\n\n    document.addEventListener('scroll', () => {\n      if (header) {\n        header.classList.toggle('shadow-sm', document.documentElement.scrollTop > 0);\n      }\n    });\n  </script><!-- Plugins and scripts required by this view--><script src=\"vendors/@coreui/utils/js/index.js\"></script><script>\n  </script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
