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

export namespace main {
	
	export class Device {
	    name: string;
	    ip: string;
	    port: string;
	
	    static createFrom(source: any = {}) {
	        return new Device(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.ip = source["ip"];
	        this.port = source["port"];
	    }
	}

}

