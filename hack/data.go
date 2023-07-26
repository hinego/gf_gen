package main

var admin = API{
	Name: "user",
	Data: []Version{
		{
			Name: "v1",
			Data: []File{
				{
					Name: "auth",
					Data: []*Function{
						{
							Name: "info",
						},
						{
							Name: "register",
						},
						{
							Name: "reset",
						},
						{
							Name: "logout",
						},
					},
				},
			},
		},
	},
}
var admin2 = API{
	Name: "data",
	Data: []Version{
		{
			Name: "v1",
			Data: []File{
				{
					Name: "auth",
					Data: []*Function{
						{
							Name: "info",
						},
						{
							Name: "register",
						},
						{
							Name: "reset",
						},
						{
							Name: "logout",
						},
					},
				},
			},
		},
		{
			Name: "v2",
			Data: []File{
				{
					Name: "auth",
					Data: []*Function{
						{
							Name: "info",
						},
						{
							Name: "register",
						},
						{
							Name: "reset",
						},
						{
							Name: "logout",
						},
					},
				},
				{
					Name: "user",
					Data: []*Function{
						{
							Name: "info",
						},
						{
							Name: "register",
						},
						{
							Name: "reset",
						},
						{
							Name: "logout",
						},
					},
				},
			},
		},
	},
}
