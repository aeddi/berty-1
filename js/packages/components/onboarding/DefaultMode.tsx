import LottieView from 'lottie-react-native'
import { useTranslation } from 'react-i18next'
import React from 'react'
import { ActivityIndicator, StatusBar, Text, View } from 'react-native'

import { ScreenFC, useNavigation } from '@berty-tech/navigation'
import {
	accountService,
	useMessengerContext,
	useNotificationsInhibitor,
	useThemeColor,
} from '@berty-tech/store'
import { useStyles } from '@berty-tech/styles'

import OnboardingWrapper from './OnboardingWrapper'
import { Icon } from '@ui-kitten/components'
import { TouchableOpacity } from 'react-native-gesture-handler'

const DefaultModeBody: React.FC = () => {
	const ctx = useMessengerContext()
	const { goBack } = useNavigation()
	const colors = useThemeColor()
	const [{ padding, border, margin, text }, { scaleSize }] = useStyles()
	const [isPressed, setIsPressed] = React.useState<boolean>(false)
	const { t }: { t: any } = useTranslation()

	const onPress = React.useCallback(async () => {
		// with an empty accountId the function returns default config
		const defaultConfig = await accountService.networkConfigGet({ accountId: '' })
		if (defaultConfig.currentConfig) {
			console.log('configForPreset', defaultConfig.currentConfig)
			setIsPressed(true)
			await ctx.createNewAccount(defaultConfig.currentConfig)
		}
	}, [ctx])

	return (
		<View style={[{ flex: 1 }]}>
			<LottieView
				source={require('./Berty_onboard_animation_assets2/Startup animation assets/Berty BG.json')}
				autoPlay
				loop
				style={{ width: '100%', position: 'absolute' }}
			/>
			{isPressed ? (
				<LottieView
					source={require('./Berty_onboard_animation_assets2/Startup animation assets/Shield dissapear.json')}
					autoPlay
					loop={false}
					style={{ position: 'absolute', top: -20, width: '100%' }}
				/>
			) : (
				<LottieView
					source={require('./Berty_onboard_animation_assets2/Startup animation assets/Shield appear.json')}
					autoPlay
					loop={false}
					style={{ position: 'absolute', top: -20, width: '100%' }}
				/>
			)}
			<View
				style={[
					padding.horizontal.medium,
					{ flex: 1, top: -(30 * scaleSize), justifyContent: 'flex-end' },
				]}
			>
				<View
					style={[
						border.shadow.large,
						padding.medium,
						border.radius.medium,
						{ backgroundColor: colors['main-background'], shadowColor: colors.shadow },
					]}
				>
					<View style={{ flexDirection: 'row', alignItems: 'center', justifyContent: 'center' }}>
						<Icon
							style={[margin.right.small]}
							name='info'
							pack='feather'
							width={23}
							fill={colors['background-header']}
						/>
						<Text
							style={[
								{
									fontFamily: 'Open Sans',
									fontWeight: '700',
									color: colors['background-header'],
									fontSize: 24 * scaleSize,
								},
							]}
						>
							{t('onboarding.default-mode.summary.title')}
						</Text>
					</View>
					<View style={[margin.top.medium]}>
						<Text
							style={[
								text.size.medium,
								{
									textAlign: 'center',
									fontFamily: 'Open Sans',
									fontWeight: '600',
									color: colors['main-text'],
								},
							]}
						>
							{t('onboarding.default-mode.summary.subtitle')}
						</Text>
					</View>
					<View style={[margin.top.medium]}>
						<Text
							style={[
								text.size.medium,
								{ fontFamily: 'Open Sans', textAlign: 'center', color: colors['main-text'] },
							]}
						>
							{t('onboarding.default-mode.summary.first-point')}
						</Text>
					</View>
					<View style={[margin.top.medium]}>
						<Text
							style={[
								text.size.medium,
								{ fontFamily: 'Open Sans', textAlign: 'center', color: colors['main-text'] },
							]}
						>
							{t('onboarding.default-mode.summary.second-point')}
						</Text>
					</View>
				</View>
				<View
					style={[
						margin.top.small,
						margin.bottom.medium,
						{
							flexDirection: 'row',
							alignItems: 'center',
							justifyContent: 'space-between',
						},
					]}
				>
					<TouchableOpacity
						style={[
							padding.medium,
							border.radius.medium,
							{ width: 170 * scaleSize, backgroundColor: '#EBECFF' },
						]}
						onPress={() => goBack()}
					>
						<Text
							style={[
								text.size.medium,
								{
									textTransform: 'uppercase',
									color: colors['background-header'],
									fontFamily: 'Open Sans',
									fontWeight: '700',
									textAlign: 'center',
								},
							]}
						>
							{t('onboarding.default-mode.summary.back-button')}
						</Text>
					</TouchableOpacity>
					<TouchableOpacity
						style={[
							padding.medium,
							border.radius.medium,
							{ width: 170 * scaleSize, backgroundColor: '#3744DD' },
						]}
						onPress={onPress}
					>
						{isPressed ? (
							<ActivityIndicator color={colors['reverted-main-text']} />
						) : (
							<Text
								style={[
									text.size.medium,
									{
										textTransform: 'uppercase',
										color: colors['reverted-main-text'],
										fontFamily: 'Open Sans',
										fontWeight: '700',
										textAlign: 'center',
									},
								]}
							>
								{t('onboarding.default-mode.summary.accept-button')}
							</Text>
						)}
					</TouchableOpacity>
				</View>
			</View>
		</View>
	)
}

export const DefaultMode: ScreenFC<'Onboarding.DefaultMode'> = () => {
	useNotificationsInhibitor(() => true)
	const colors = useThemeColor()

	return (
		<OnboardingWrapper>
			<StatusBar backgroundColor={colors['background-header']} barStyle='light-content' />
			<DefaultModeBody />
		</OnboardingWrapper>
	)
}
