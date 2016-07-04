package rubiksimg

import (
	"image/color"

	"github.com/llgcode/draw2d"
)

func drawRightFace(ctx draw2d.GraphicContext, scale float64, colors []color.Color) {
	ctx.SetFillColor(colors[0])
	ctx.MoveTo(scale*13.305428, scale*51.540590)
	ctx.LineTo(scale*40.343717, scale*72.760901)
	ctx.CubicCurveTo(scale*43.734039, scale*75.464962, scale*45.237360, scale*82.490187,
		scale*45.767910, scale*86.495547)
	ctx.LineTo(scale*49.071660, scale*165.667350)
	ctx.CubicCurveTo(scale*49.289247, scale*172.150290, scale*47.781718, scale*176.557160,
		scale*44.096549, scale*172.925110)
	ctx.LineTo(scale*14.319497, scale*148.460780)
	ctx.CubicCurveTo(scale*12.537448, scale*146.848450, scale*11.649146, scale*144.480610,
		scale*11.278686, scale*142.076270)
	ctx.LineTo(scale*6.684044, scale*55.738325)
	ctx.CubicCurveTo(scale*5.914454, scale*51.334595, scale*9.148533, scale*49.897364,
		scale*13.305428, scale*51.540590)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[1])
	ctx.MoveTo(scale*60.925762, scale*84.484801)
	ctx.LineTo(scale*94.440727, scale*110.173380)
	ctx.CubicCurveTo(scale*98.484160, scale*113.287130, scale*98.117046, scale*115.753300,
		scale*98.906924, scale*122.400750)
	ctx.LineTo(scale*99.906265, scale*206.016210)
	ctx.CubicCurveTo(scale*100.082290, scale*212.185990, scale*98.954748, scale*216.917160,
		scale*94.922195, scale*215.058640)
	ctx.LineTo(scale*59.491492, scale*185.322390)
	ctx.CubicCurveTo(scale*55.399174, scale*181.448720, scale*54.141798, scale*179.024610,
		scale*53.773149, scale*173.562040)
	ctx.LineTo(scale*51.281758, scale*92.445638)
	ctx.CubicCurveTo(scale*50.977610, scale*83.275467, scale*56.920401, scale*82.791038,
		scale*60.925762, scale*84.484801)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[2])
	ctx.MoveTo(scale*114.953940, scale*121.045210)
	ctx.LineTo(scale*159.771560, scale*153.233820)
	ctx.CubicCurveTo(scale*163.559180, scale*156.470830, scale*165.579320, scale*161.453600,
		scale*165.851460, scale*169.729030)
	ctx.LineTo(scale*165.123240, scale*258.384880)
	ctx.CubicCurveTo(scale*165.143180, scale*269.270490, scale*163.081960, scale*270.870600,
		scale*159.695010, scale*268.501700)
	ctx.LineTo(scale*113.597590, scale*227.818200)
	ctx.CubicCurveTo(scale*108.062330, scale*222.322250, scale*108.218040, scale*221.408810,
		scale*107.697890, scale*213.179000)
	ctx.LineTo(scale*106.780440, scale*131.955830)
	ctx.CubicCurveTo(scale*106.253030, scale*119.343540, scale*109.796840, scale*118.294520,
		scale*114.953940, scale*121.045210)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[3])
	ctx.MoveTo(scale*15.880212, scale*156.133970)
	ctx.LineTo(scale*45.353913, scale*181.289430)
	ctx.CubicCurveTo(scale*50.303506, scale*186.196870, scale*51.122928, scale*188.031090,
		scale*51.146692, scale*194.603600)
	ctx.LineTo(scale*55.126981, scale*267.337310)
	ctx.CubicCurveTo(scale*55.038091, scale*274.160790, scale*54.210515, scale*276.038730,
		scale*50.685456, scale*273.367280)
	ctx.LineTo(scale*20.772099, scale*242.788520)
	ctx.CubicCurveTo(scale*18.215174, scale*239.755780, scale*17.772229, scale*237.411820,
		scale*17.401769, scale*232.926170)
	ctx.LineTo(scale*12.455448, scale*161.175690)
	ctx.CubicCurveTo(scale*11.924888, scale*156.529920, scale*13.177442, scale*153.868160,
		scale*15.880212, scale*156.133970)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[4])
	ctx.MoveTo(scale*63.560739, scale*196.668590)
	ctx.LineTo(scale*95.849604, scale*224.215480)
	ctx.CubicCurveTo(scale*99.860002, scale*227.900360, scale*101.420680, scale*227.848830,
		scale*102.248290, scale*236.168620)
	ctx.LineTo(scale*105.161150, scale*322.881310)
	ctx.CubicCurveTo(scale*105.329610, scale*328.834720, scale*100.635400, scale*326.982890,
		scale*97.270450, scale*323.991240)
	ctx.LineTo(scale*64.004863, scale*289.914880)
	ctx.CubicCurveTo(scale*60.376668, scale*285.922340, scale*60.812901, scale*283.086500,
		scale*60.282344, scale*279.081140)
	ctx.LineTo(scale*56.974857, scale*203.970330)
	ctx.CubicCurveTo(scale*56.534832, scale*197.334610, scale*59.555379, scale*193.983720,
		scale*63.560739, scale*196.668590)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[5])
	ctx.MoveTo(scale*115.165350, scale*240.136900)
	ctx.LineTo(scale*157.390730, scale*278.017180)
	ctx.CubicCurveTo(scale*163.353530, scale*283.617540, scale*164.201130, scale*287.753760,
		scale*164.337130, scale*292.739150)
	ctx.LineTo(scale*164.464830, scale*372.926510)
	ctx.CubicCurveTo(scale*164.837100, scale*382.725590, scale*163.118730, scale*386.721700,
		scale*158.318310, scale*383.060650)
	ctx.LineTo(scale*116.216280, scale*341.592770)
	ctx.CubicCurveTo(scale*113.144360, scale*337.418810, scale*111.772190, scale*331.794540,
		scale*111.565450, scale*326.252280)
	ctx.LineTo(scale*108.059310, scale*249.470500)
	ctx.CubicCurveTo(scale*107.834300, scale*242.185660, scale*110.331410, scale*236.363050,
		scale*115.165350, scale*240.136900)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[6])
	ctx.MoveTo(scale*21.801396, scale*250.847880)
	ctx.LineTo(scale*53.362682, scale*285.050810)
	ctx.CubicCurveTo(scale*55.478478, scale*286.947370, scale*55.860452, scale*293.075040,
		scale*56.230911, scale*297.080390)
	ctx.LineTo(scale*59.037291, scale*362.099420)
	ctx.CubicCurveTo(scale*59.567851, scale*365.304280, scale*57.391745, scale*365.057950,
		scale*55.670483, scale*362.854910)
	ctx.LineTo(scale*26.987682, scale*329.121490)
	ctx.CubicCurveTo(scale*24.403644, scale*326.580000, scale*22.851374, scale*322.172070,
		scale*22.801115, scale*318.166720)
	ctx.LineTo(scale*18.702227, scale*254.063490)
	ctx.CubicCurveTo(scale*18.171667, scale*250.058130, scale*18.780430, scale*246.996130,
		scale*21.801396, scale*250.847880)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[7])
	ctx.MoveTo(scale*63.584727, scale*297.864820)
	ctx.LineTo(scale*100.097710, scale*336.258750)
	ctx.CubicCurveTo(scale*103.855320, scale*339.948770, scale*105.542250, scale*346.524610,
		scale*105.619970, scale*350.869590)
	ctx.LineTo(scale*106.700760, scale*413.699910)
	ctx.CubicCurveTo(scale*106.896610, scale*418.854270, scale*106.324150, scale*421.593830,
		scale*102.848520, scale*417.566110)
	ctx.LineTo(scale*67.414737, scale*377.504340)
	ctx.CubicCurveTo(scale*64.508880, scale*374.540930, scale*64.456140, scale*374.604960,
		scale*63.999005, scale*369.270860)
	ctx.LineTo(scale*60.477610, scale*304.569610)
	ctx.CubicCurveTo(scale*60.281747, scale*300.540090, scale*59.899568, scale*294.232770,
		scale*63.584727, scale*297.864820)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[8])
	ctx.MoveTo(scale*118.860810, scale*353.358710)
	ctx.LineTo(scale*159.480130, scale*395.623800)
	ctx.CubicCurveTo(scale*162.517640, scale*399.133800, scale*164.424500, scale*404.651780,
		scale*165.106100, scale*412.626360)
	ctx.LineTo(scale*166.135870, scale*476.201080)
	ctx.CubicCurveTo(scale*166.979380, scale*485.118150, scale*165.130280, scale*489.476540,
		scale*161.923610, scale*484.858520)
	ctx.LineTo(scale*118.272860, scale*435.659510)
	ctx.CubicCurveTo(scale*114.349710, scale*431.594940, scale*113.966820, scale*428.882850,
		scale*114.076660, scale*422.636080)
	ctx.LineTo(scale*112.712210, scale*363.084690)
	ctx.CubicCurveTo(scale*112.074730, scale*354.599640, scale*115.240410, scale*350.550980,
		scale*118.860810, scale*353.358710)
	ctx.Close()
	ctx.Fill()
}

func drawFrontFace(ctx draw2d.GraphicContext, scale float64, colors []color.Color) {
	ctx.SetFillColor(colors[0])
	ctx.MoveTo(scale*182.178510, scale*158.388410)
	ctx.LineTo(scale*273.788920, scale*136.410740)
	ctx.CubicCurveTo(scale*277.794270, scale*135.880180, scale*280.165130, scale*138.037180,
		scale*280.695680, scale*142.042540)
	ctx.LineTo(scale*275.834450, scale*238.474950)
	ctx.CubicCurveTo(scale*275.564510, scale*242.320210, scale*270.928300, scale*248.688840,
		scale*266.922950, scale*250.340110)
	ctx.LineTo(scale*181.549290, scale*272.707500)
	ctx.CubicCurveTo(scale*177.543930, scale*273.238060, scale*174.372580, scale*267.403470,
		scale*174.322330, scale*263.398110)
	ctx.LineTo(scale*174.451510, scale*169.729120)
	ctx.CubicCurveTo(scale*174.881550, scale*165.563670, scale*178.333260, scale*160.039680,
		scale*182.178510, scale*158.388410)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[1])
	ctx.MoveTo(scale*301.533600, scale*130.911780)
	ctx.LineTo(scale*381.445870, scale*113.649400)
	ctx.CubicCurveTo(scale*386.246530, scale*112.772630, scale*388.765240, scale*113.368590,
		scale*388.005160, scale*120.526770)
	ctx.LineTo(scale*379.365520, scale*210.847340)
	ctx.CubicCurveTo(scale*378.615280, scale*214.532500, scale*378.836500, scale*220.530400,
		scale*371.629130, scale*221.701360)
	ctx.LineTo(scale*293.604020, scale*243.457870)
	ctx.CubicCurveTo(scale*289.598660, scale*243.988420, scale*285.326310, scale*241.085870,
		scale*284.795750, scale*237.080520)
	ctx.LineTo(scale*289.940140, scale*142.544300)
	ctx.CubicCurveTo(scale*290.137030, scale*136.892880, scale*296.598420, scale*132.556810,
		scale*301.533600, scale*130.911780)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[2])
	ctx.MoveTo(scale*406.907450, scale*110.062680)
	ctx.LineTo(scale*473.534830, scale*95.080791)
	ctx.CubicCurveTo(scale*479.983350, scale*94.205963, scale*478.799760, scale*98.364434,
		scale*478.087960, scale*104.194870)
	ctx.LineTo(scale*467.374240, scale*185.580620)
	ctx.CubicCurveTo(scale*466.948710, scale*190.262530, scale*463.678950, scale*194.011820,
		scale*459.513500, scale*195.663080)
	ctx.LineTo(scale*398.568880, scale*214.054370)
	ctx.CubicCurveTo(scale*391.807150, scale*215.885320, scale*389.296970, scale*209.675070,
		scale*389.566920, scale*205.189420)
	ctx.LineTo(scale*398.495680, scale*119.713910)
	ctx.CubicCurveTo(scale*398.925730, scale*115.708560, scale*402.902090, scale*111.951730,
		scale*406.907450, scale*110.062680)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[3])
	ctx.MoveTo(scale*180.403970, scale*283.860000)
	ctx.LineTo(scale*263.366350, scale*261.105010)
	ctx.CubicCurveTo(scale*267.371700, scale*260.414350, scale*274.529440, scale*262.662370,
		scale*274.090350, scale*269.271530)
	ctx.LineTo(scale*270.061710, scale*352.515730)
	ctx.CubicCurveTo(scale*269.630150, scale*358.628000, scale*262.098910, scale*361.822870,
		scale*258.413760, scale*363.314040)
	ctx.LineTo(scale*181.857750, scale*388.147890)
	ctx.CubicCurveTo(scale*177.852390, scale*388.678450, scale*173.671550, scale*388.186410,
		scale*173.133220, scale*383.341780)
	ctx.LineTo(scale*173.635250, scale*298.803250)
	ctx.CubicCurveTo(scale*173.584990, scale*295.598380, scale*175.906260, scale*285.671360,
		scale*180.403970, scale*283.860000)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[4])
	ctx.MoveTo(scale*292.972170, scale*253.393720)
	ctx.LineTo(scale*369.040340, scale*231.810550)
	ctx.CubicCurveTo(scale*373.045690, scale*231.280000, scale*377.680550, scale*234.108220,
		scale*377.557150, scale*240.249120)
	ctx.LineTo(scale*370.219790, scale*316.723460)
	ctx.CubicCurveTo(scale*369.949840, scale*321.529330, scale*364.677580, scale*326.480350,
		scale*361.622270, scale*327.497240)
	ctx.LineTo(scale*288.849550, scale*352.520860)
	ctx.CubicCurveTo(scale*284.644920, scale*354.019550, scale*280.273320, scale*348.357450,
		scale*279.742760, scale*344.352090)
	ctx.LineTo(scale*283.597730, scale*265.296000)
	ctx.CubicCurveTo(scale*283.707570, scale*260.810350, scale*287.829540, scale*255.205080,
		scale*292.972170, scale*253.393720)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[5])
	ctx.MoveTo(scale*393.257250, scale*226.028330)
	ctx.LineTo(scale*459.130590, scale*204.762730)
	ctx.CubicCurveTo(scale*463.830180, scale*203.400880, scale*464.561690, scale*207.213150,
		scale*464.131650, scale*211.218510)
	ctx.LineTo(scale*453.012820, scale*288.384290)
	ctx.CubicCurveTo(scale*452.665860, scale*292.487700, scale*451.548370, scale*293.973060,
		scale*447.962620, scale*295.542070)
	ctx.LineTo(scale*383.523150, scale*317.937580)
	ctx.CubicCurveTo(scale*379.272390, scale*318.670880, scale*377.827440, scale*315.821090,
		scale*378.333770, scale*310.715990)
	ctx.LineTo(scale*385.766580, scale*235.103530)
	ctx.CubicCurveTo(scale*386.050770, scale*232.713450, scale*388.879710, scale*227.099910,
		scale*393.257250, scale*226.028330)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[6])
	ctx.MoveTo(scale*181.897300, scale*397.299240)
	ctx.LineTo(scale*261.701940, scale*370.127280)
	ctx.CubicCurveTo(scale*271.504480, scale*367.299980, scale*270.640000, scale*373.527480,
		scale*270.370050, scale*377.532830)
	ctx.LineTo(scale*266.827380, scale*448.045760)
	ctx.CubicCurveTo(scale*266.877640, scale*452.051120, scale*265.061810, scale*457.310500,
		scale*261.056460, scale*459.442060)
	ctx.LineTo(scale*183.587360, scale*490.340340)
	ctx.CubicCurveTo(scale*179.582010, scale*490.870900, scale*175.108890, scale*488.073500,
		scale*174.578340, scale*484.068150)
	ctx.LineTo(scale*174.447500, scale*410.435270)
	ctx.CubicCurveTo(scale*174.717440, scale*406.429910, scale*176.438420, scale*399.924520,
		scale*181.897300, scale*397.299240)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[7])
	ctx.MoveTo(scale*285.958380, scale*361.603590)
	ctx.LineTo(scale*360.964460, scale*334.706510)
	ctx.CubicCurveTo(scale*365.129910, scale*333.055250, scale*369.623030, scale*334.335740,
		scale*369.249140, scale*340.303390)
	ctx.LineTo(scale*360.912460, scale*412.001850)
	ctx.CubicCurveTo(scale*360.753990, scale*416.528590, scale*359.691400, scale*419.425570,
		scale*356.429150, scale*421.211710)
	ctx.LineTo(scale*282.171640, scale*450.931830)
	ctx.CubicCurveTo(scale*278.166280, scale*451.462390, scale*275.118220, scale*448.783710,
		scale*275.388160, scale*444.618250)
	ctx.LineTo(scale*279.031300, scale*374.050680)
	ctx.CubicCurveTo(scale*279.002000, scale*368.160110, scale*281.953020, scale*363.254860,
		scale*285.958380, scale*361.603590)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[8])
	ctx.MoveTo(scale*382.515920, scale*326.462710)
	ctx.LineTo(scale*444.081060, scale*304.853330)
	ctx.CubicCurveTo(scale*449.367210, scale*303.041970, scale*451.119410, scale*305.024380,
		scale*450.529270, scale*309.510040)
	ctx.LineTo(scale*439.534300, scale*381.872980)
	ctx.CubicCurveTo(scale*438.198590, scale*385.066460, scale*437.802190, scale*387.370910,
		scale*434.117040, scale*388.823220)
	ctx.LineTo(scale*372.466700, scale*414.121590)
	ctx.CubicCurveTo(scale*367.749050, scale*415.726950, scale*367.138130, scale*411.964230,
		scale*367.474390, scale*407.716370)
	ctx.LineTo(scale*375.652860, scale*337.864080)
	ctx.CubicCurveTo(scale*376.082910, scale*333.058230, scale*378.787280, scale*327.873810,
		scale*382.515920, scale*326.462710)
	ctx.Close()
	ctx.Fill()
}

func drawTopFace(ctx draw2d.GraphicContext, scale float64, colors []color.Color) {
	ctx.SetFillColor(colors[0])
	ctx.MoveTo(scale*208.015900, scale*16.168793)
	ctx.LineTo(scale*269.753340, scale*6.060053)
	ctx.CubicCurveTo(scale*273.758700, scale*5.529493, scale*275.214650, scale*4.917190,
		scale*278.801420, scale*6.793998)
	ctx.LineTo(scale*322.610240, scale*24.147523)
	ctx.CubicCurveTo(scale*325.358840, scale*25.591285, scale*326.141010, scale*26.455624,
		scale*322.135640, scale*27.626585)
	ctx.LineTo(scale*261.396360, scale*37.422514)
	ctx.CubicCurveTo(scale*257.043090, scale*37.934465, scale*249.375710, scale*37.324805,
		scale*245.624890, scale*35.579744)
	ctx.LineTo(scale*207.698020, scale*19.306437)
	ctx.CubicCurveTo(scale*204.242400, scale*18.022573, scale*205.124320, scale*16.294930,
		scale*208.015900, scale*16.168793)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[1])
	ctx.MoveTo(scale*270.311910, scale*40.717976)
	ctx.LineTo(scale*328.466220, scale*30.954512)
	ctx.CubicCurveTo(scale*331.510970, scale*29.783553, scale*335.471800, scale*29.045434,
		scale*339.489370, scale*31.094387)
	ctx.LineTo(scale*389.781750, scale*51.331875)
	ctx.CubicCurveTo(scale*394.042420, scale*53.506556, scale*393.023900, scale*54.981242,
		scale*389.497210, scale*55.365944)
	ctx.LineTo(scale*326.934390, scale*66.861350)
	ctx.CubicCurveTo(scale*322.768930, scale*67.712108, scale*315.780670, scale*66.433592,
		scale*312.688510, scale*64.669634)
	ctx.LineTo(scale*269.272130, scale*44.479849)
	ctx.CubicCurveTo(scale*266.461000, scale*43.031581, scale*266.626750, scale*41.728835,
		scale*270.311910, scale*40.717976)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[2])
	ctx.MoveTo(scale*336.566960, scale*69.287076)
	ctx.LineTo(scale*397.636690, scale*58.137148)
	ctx.CubicCurveTo(scale*401.086410, scale*57.208951, scale*404.778250, scale*58.130822,
		scale*410.425770, scale*60.353676)
	ctx.LineTo(scale*464.508550, scale*81.433424)
	ctx.CubicCurveTo(scale*470.529060, scale*83.521457, scale*473.468440, scale*87.244831,
		scale*469.463070, scale*89.216293)
	ctx.LineTo(scale*411.530760, scale*101.470450)
	ctx.CubicCurveTo(scale*408.147380, scale*102.502900, scale*397.189320, scale*102.662560,
		scale*394.099790, scale*100.667720)
	ctx.LineTo(scale*336.105750, scale*74.185386)
	ctx.CubicCurveTo(scale*332.764730, scale*71.922386, scale*332.666750, scale*69.986992,
		scale*336.566960, scale*69.287076)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[3])
	ctx.MoveTo(scale*116.967040, scale*29.014065)
	ctx.LineTo(scale*190.193710, scale*18.227759)
	ctx.CubicCurveTo(scale*193.076740, scale*17.920153, scale*196.606750, scale*18.927933,
		scale*197.950650, scale*19.968175)
	ctx.LineTo(scale*231.617500, scale*34.951899)
	ctx.CubicCurveTo(scale*234.874040, scale*36.701617, scale*234.418900, scale*38.896015,
		scale*231.534230, scale*39.586675)
	ctx.LineTo(scale*157.735500, scale*51.255524)
	ctx.CubicCurveTo(scale*153.239700, scale*51.961389, scale*146.989730, scale*51.849221,
		scale*143.961530, scale*49.794623)
	ctx.LineTo(scale*112.967770, scale*33.044845)
	ctx.CubicCurveTo(scale*109.396120, scale*31.009309, scale*112.961690, scale*29.864816,
		scale*116.967040, scale*29.014065)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[4])
	ctx.MoveTo(scale*170.482790, scale*53.739436)
	ctx.LineTo(scale*241.651150, scale*41.973692)
	ctx.CubicCurveTo(scale*245.656510, scale*41.443135, scale*251.717640, scale*41.248700,
		scale*255.130010, scale*43.492950)
	ctx.LineTo(scale*298.715100, scale*63.776222)
	ctx.CubicCurveTo(scale*302.287560, scale*65.700282, scale*303.169380, scale*68.298399,
		scale*299.164020, scale*68.828956)
	ctx.LineTo(scale*221.423960, scale*83.385817)
	ctx.CubicCurveTo(scale*218.058990, scale*83.916375, scale*210.913850, scale*82.856747,
		scale*207.821680, scale*81.092790)
	ctx.LineTo(scale*166.350150, scale*58.667629)
	ctx.CubicCurveTo(scale*162.777690, scale*56.103179, scale*166.477420, scale*54.269994,
		scale*170.482790, scale*53.739436)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[5])
	ctx.MoveTo(scale*234.178580, scale*86.640558)
	ctx.LineTo(scale*308.088400, scale*72.881494)
	ctx.CubicCurveTo(scale*312.093770, scale*72.350935, scale*318.626430, scale*71.610297,
		scale*322.839300, scale*74.494959)
	ctx.LineTo(scale*380.232970, scale*101.016140)
	ctx.CubicCurveTo(scale*382.844840, scale*103.260380, scale*383.275890, scale*107.177310,
		scale*379.270530, scale*108.508370)
	ctx.LineTo(scale*300.249640, scale*123.424200)
	ctx.CubicCurveTo(scale*296.244280, scale*124.755250, scale*290.329430, scale*123.925920,
		scale*285.476160, scale*120.881160)
	ctx.LineTo(scale*232.466880, scale*92.266536)
	ctx.CubicCurveTo(scale*229.374710, scale*90.342477, scale*230.173220, scale*87.971619,
		scale*234.178580, scale*86.640558)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[6])
	ctx.MoveTo(scale*23.348625, scale*43.167759)
	ctx.LineTo(scale*91.061974, scale*32.897828)
	ctx.CubicCurveTo(scale*95.009786, scale*32.418498, scale*99.005619, scale*31.916320,
		scale*101.357050, scale*33.680264)
	ctx.LineTo(scale*132.523360, scale*50.748834)
	ctx.CubicCurveTo(scale*137.363900, scale*53.755196, scale*135.577590, scale*56.641768,
		scale*131.572240, scale*57.172328)
	ctx.LineTo(scale*56.551277, scale*69.765807)
	ctx.CubicCurveTo(scale*51.992688, scale*70.353383, scale*48.930844, scale*70.408699,
		scale*46.049907, scale*68.279732)
	ctx.LineTo(scale*15.629868, scale*47.347863)
	ctx.CubicCurveTo(scale*12.365004, scale*45.277954, scale*20.409636, scale*43.572566,
		scale*23.348625, scale*43.167759)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[7])
	ctx.MoveTo(scale*67.927185, scale*70.656234)
	ctx.LineTo(scale*142.847190, scale*58.313290)
	ctx.CubicCurveTo(scale*146.852540, scale*57.782730, scale*155.169860, scale*58.670511,
		scale*158.902420, scale*60.754655)
	ctx.LineTo(scale*197.741240, scale*82.476005)
	ctx.CubicCurveTo(scale*200.353110, scale*84.720260, scale*199.140530, scale*86.204014,
		scale*195.615490, scale*87.535076)
	ctx.LineTo(scale*113.382320, scale*104.583500)
	ctx.CubicCurveTo(scale*109.376960, scale*105.594360, scale*103.198400, scale*105.494960,
		scale*99.946129, scale*103.250710)
	ctx.LineTo(scale*63.593588, scale*77.511289)
	ctx.CubicCurveTo(scale*61.301923, scale*75.267034, scale*63.921825, scale*71.506995,
		scale*67.927185, scale*70.656234)
	ctx.Close()
	ctx.Fill()

	ctx.SetFillColor(colors[8])
	ctx.MoveTo(scale*123.162040, scale*107.417340)
	ctx.LineTo(scale*203.880510, scale*90.359284)
	ctx.CubicCurveTo(scale*207.885880, scale*89.828726, scale*217.641160, scale*90.981557,
		scale*220.733320, scale*93.385916)
	ctx.LineTo(scale*270.901830, scale*120.441840)
	ctx.CubicCurveTo(scale*273.993990, scale*122.846190, scale*276.264860, scale*128.629380,
		scale*272.099400, scale*130.280640)
	ctx.LineTo(scale*183.265110, scale*149.717790)
	ctx.CubicCurveTo(scale*178.619340, scale*151.164820, scale*171.278720, scale*150.772340,
		scale*167.225960, scale*147.475910)
	ctx.LineTo(scale*119.639010, scale*114.495590)
	ctx.CubicCurveTo(scale*116.066540, scale*112.251330, scale*119.316790, scale*108.588300,
		scale*123.162040, scale*107.417340)
	ctx.Close()
	ctx.Fill()
}
