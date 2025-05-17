export namespace config {
	
	export class Configuration {
	    ScreenSize: number[];
	    Method: string;
	    HaoiUser: string;
	    HaoiRebate: string;
	    TwoCaptchaKey: string;
	
	    static createFrom(source: any = {}) {
	        return new Configuration(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ScreenSize = source["ScreenSize"];
	        this.Method = source["Method"];
	        this.HaoiUser = source["HaoiUser"];
	        this.HaoiRebate = source["HaoiRebate"];
	        this.TwoCaptchaKey = source["TwoCaptchaKey"];
	    }
	}

}

