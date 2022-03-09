import React, { useCallback, useEffect, useState } from 'react'
import { Linking } from 'react-native'
import { useTranslation } from 'react-i18next'
import {
	createNativeStackNavigator,
	NativeStackNavigationOptions,
} from '@react-navigation/native-stack'
import { CommonActions, NavigationProp, useNavigation } from '@react-navigation/native'
import { Icon } from '@ui-kitten/components'
import mapValues from 'lodash/mapValues'

import * as RawComponents from '@berty-tech/components'
import { useMessengerContext, useThemeColor } from '@berty-tech/store'
import { useStyles } from '@berty-tech/styles'

import { dispatch } from './rootRef'
import { ScreensParams } from './types'
import { useSelector } from 'react-redux'
import { MESSENGER_APP_STATE, selectAppState } from '@berty-tech/redux/reducers/ui.reducer'

export const CustomTitleStyle: () => any = () => {
	const [{ text }] = useStyles()

	return [
		text.size.large,
		{
			headerTitleStyle: {
				fontFamily: 'Open Sans',
				fontWeight: '700',
			},
		},
	]
}

const ChatScreenOptions: (
	additionalProps?: NativeStackNavigationOptions,
) => NativeStackNavigationOptions = additionalProps => {
	const colors = useThemeColor()
	return {
		headerStyle: {
			backgroundColor: colors['main-background'],
		},
		headerTintColor: colors['main-text'],
		headerBackTitleVisible: false,
		headerShadowVisible: false,
		...additionalProps,
	}
}

const BackgroundHeaderScreenOptions: (
	additionalProps?: NativeStackNavigationOptions,
) => NativeStackNavigationOptions = additionalProps => {
	const colors = useThemeColor()
	return {
		headerStyle: {
			backgroundColor: colors['background-header'],
		},
		headerTintColor: colors['reverted-main-text'],
		headerBackTitleVisible: false,
		headerShadowVisible: false,
		...additionalProps,
	}
}

const SecondaryBackgroundHeaderScreenOptions: (
	additionalProps?: NativeStackNavigationOptions,
) => NativeStackNavigationOptions = additionalProps => {
	const colors = useThemeColor()
	return {
		headerStyle: {
			backgroundColor: colors['secondary-background-header'],
		},
		headerTintColor: colors['reverted-main-text'],
		headerBackTitleVisible: false,
		headerShadowVisible: false,
		...additionalProps,
	}
}

const AltBackgroundHeaderScreenOptions: (
	additionalProps?: NativeStackNavigationOptions,
) => NativeStackNavigationOptions = additionalProps => {
	const colors = useThemeColor()
	return {
		headerStyle: {
			backgroundColor: colors['alt-secondary-background-header'],
		},
		headerTintColor: colors['reverted-main-text'],
		headerBackTitleVisible: false,
		headerShadowVisible: false,
		...additionalProps,
	}
}

function useLinking(): [string | null, unknown] {
	const [url, setUrl] = useState<string | null>(null)
	const [error, setError] = useState<unknown>()

	const initialUrl = useCallback(async () => {
		try {
			const linkingUrl = await Linking.getInitialURL()
			if (linkingUrl) {
				setUrl(linkingUrl)
			}
		} catch (ex) {
			setError(ex)
		}
	}, [])

	useEffect(() => {
		const handleOpenUrl = (ev: any) => {
			console.log('handleOpenUrl:', ev.url)
			setUrl(null)
			setUrl(ev.url)
		}

		// for initial render
		initialUrl().then(() => {
			Linking.addEventListener('url', handleOpenUrl)
		})

		return () => Linking.removeEventListener('url', handleOpenUrl)
	}, [initialUrl])

	return [url, error]
}

const DeepLinkBridge: React.FC = React.memo(() => {
	const [url, error] = useLinking()
	const navigation = useNavigation<NavigationProp<ScreensParams>>()
	const ctx = useMessengerContext()

	useEffect(() => {
		if (!ctx.handledLink && url && !error && !(url as string).startsWith('berty://services-auth')) {
			ctx.setHandledLink(true)
			navigation.navigate('Modals.ManageDeepLink', { type: 'link', value: url })
		}
	}, [ctx, error, navigation, url])

	return null
})

let Components: typeof RawComponents

// @ts-ignore
Components = mapValues(RawComponents, SubComponents =>
	mapValues(
		SubComponents,
		(Component: React.FC): React.FC =>
			React.memo(props => (
				<>
					<DeepLinkBridge />
					<Component {...props} />
				</>
			)),
	),
)

const NavigationStack = createNativeStackNavigator<ScreensParams>()

export const Navigation: React.FC = React.memo(() => {
	const appState = useSelector(selectAppState)
	const colors = useThemeColor()
	const [, { scaleSize }] = useStyles()
	const { t }: any = useTranslation()

	useEffect(() => {
		console.log('context app State', appState)
		switch (appState) {
			case MESSENGER_APP_STATE.READY:
				dispatch(
					CommonActions.reset({
						routes: [{ name: 'Main.Home' }],
					}),
				)
				return
			case MESSENGER_APP_STATE.PRE_READY:
				dispatch(
					CommonActions.reset({
						routes: [{ name: 'Onboarding.SetupFinished' }],
					}),
				)
				return
			case MESSENGER_APP_STATE.GET_STARTED:
				dispatch(
					CommonActions.reset({
						routes: [{ name: 'Onboarding.GetStarted' }],
					}),
				)
				return
		}
	}, [appState])

	return (
		<NavigationStack.Navigator
			initialRouteName={
				appState === MESSENGER_APP_STATE.GET_STARTED ? 'Onboarding.GetStarted' : 'Main.Home'
			}
		>
			{/* OnBoarding */}
			<NavigationStack.Screen
				name={'Onboarding.GetStarted'}
				component={Components.Onboarding.GetStarted}
				options={{ headerShown: false }}
			/>
			<NavigationStack.Screen
				name={'Onboarding.CreateAccount'}
				component={Components.Onboarding.CreateAccount}
				options={{
					headerStyle: {
						backgroundColor: colors['background-header'],
					},
					headerTintColor: colors['reverted-main-text'],
					headerBackTitleVisible: false,
					title: '',
				}}
			/>
			<NavigationStack.Screen
				name={'Onboarding.SetupFinished'}
				component={Components.Onboarding.SetupFinished}
				options={{ headerShown: false }}
			/>
			<NavigationStack.Screen
				name={'Onboarding.CustomModeSettings'}
				component={Components.Onboarding.CustomModeSettings}
				options={{
					headerStyle: {
						backgroundColor: colors['background-header'],
					},
					headerTintColor: colors['reverted-main-text'],
					headerBackTitleVisible: false,
					title: '',
				}}
			/>
			<NavigationStack.Screen
				name={'Onboarding.WebViews'}
				component={Components.Onboarding.WebViews}
				options={{ title: '', headerBackTitle: '', headerTintColor: colors['main-text'] }}
			/>
			<NavigationStack.Screen
				name={'Onboarding.DefaultMode'}
				component={Components.Onboarding.DefaultMode}
				options={{
					headerStyle: {
						backgroundColor: colors['background-header'],
					},
					headerTintColor: colors['reverted-main-text'],
					headerBackTitleVisible: false,
					title: '',
				}}
			/>
			<NavigationStack.Screen
				name={'Onboarding.CustomMode'}
				component={Components.Onboarding.CustomMode}
				options={{
					headerStyle: {
						backgroundColor: colors['background-header'],
					},
					headerTintColor: colors['reverted-main-text'],
					headerBackTitleVisible: false,
					title: '',
				}}
			/>
			{/* Main */}
			<NavigationStack.Screen
				name={'Main.Home'}
				component={Components.Main.Home}
				options={{ headerShown: false }}
			/>
			<NavigationStack.Screen
				name={'Main.ContactRequest'}
				component={Components.Main.ContactRequest}
			/>
			<NavigationStack.Screen
				name={'Main.Scan'}
				component={Components.Main.Scan}
				options={SecondaryBackgroundHeaderScreenOptions({
					title: t('main.scan.title'),
					headerRight: () => (
						<Icon
							name='qr'
							pack='custom'
							width={35 * scaleSize}
							height={35 * scaleSize}
							fill={colors['reverted-main-text']}
						/>
					),
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Main.Permissions'}
				component={Components.Main.Permissions}
				options={{ headerShown: false, presentation: 'formSheet' }}
			/>
			<NavigationStack.Screen
				name={'Main.BlePermission'}
				component={Components.Main.BlePermission}
				options={{ headerShown: false, presentation: 'formSheet' }}
			/>
			{/* CreateGroup */}
			<NavigationStack.Screen
				name={'Main.CreateGroupAddMembers'}
				component={Components.Main.CreateGroupAddMembers}
				options={BackgroundHeaderScreenOptions({
					title: t('main.home.create-group.title'),
					headerRight: () => (
						<Icon
							name='users'
							pack='custom'
							width={35 * scaleSize}
							height={35 * scaleSize}
							fill={colors['reverted-main-text']}
						/>
					),
					...CustomTitleStyle(),
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Main.CreateGroupFinalize'}
				component={Components.Main.CreateGroupFinalize}
				options={BackgroundHeaderScreenOptions({
					title: t('main.home.create-group.title'),
					headerRight: () => (
						<Icon
							name='users'
							pack='custom'
							width={35 * scaleSize}
							height={35 * scaleSize}
							fill={colors['reverted-main-text']}
						/>
					),
					...CustomTitleStyle(),
					presentation: 'formSheet',
				})}
			/>
			{/* Chat */}
			<NavigationStack.Screen
				name={'Chat.OneToOne'}
				component={Components.Chat.OneToOne}
				options={ChatScreenOptions({
					...CustomTitleStyle(),
				})}
			/>
			<NavigationStack.Screen
				name={'Chat.Group'}
				component={Components.Chat.MultiMember}
				options={ChatScreenOptions({
					...ChatScreenOptions(),
				})}
			/>
			<NavigationStack.Screen
				name={'Chat.OneToOneSettings'}
				component={Components.Chat.OneToOneSettings}
				options={BackgroundHeaderScreenOptions({
					title: '',
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Chat.ContactSettings'}
				component={Components.Chat.ContactSettings}
				options={BackgroundHeaderScreenOptions({
					title: '',
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Group.MultiMemberSettings'}
				component={Components.Chat.MultiMemberSettings}
				options={BackgroundHeaderScreenOptions({
					title: '',
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Group.ChatSettingsMemberDetail'}
				component={Components.Chat.ChatSettingsMemberDetail}
				options={BackgroundHeaderScreenOptions({
					title: '',
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Chat.MultiMemberQR'}
				component={Components.Chat.MultiMemberQR}
				options={BackgroundHeaderScreenOptions({
					title: t('chat.multi-member-qr.title'),
					...CustomTitleStyle(),
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Group.MultiMemberSettingsAddMembers'}
				component={Components.Chat.MultiMemberSettingsAddMembers}
				options={BackgroundHeaderScreenOptions({
					title: t('chat.add-members.members'),
					...CustomTitleStyle(),
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Chat.ReplicateGroupSettings'}
				component={Components.Chat.ReplicateGroupSettings}
				options={BackgroundHeaderScreenOptions({
					title: t('chat.replicate-group-settings.title'),
					...CustomTitleStyle(),
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Chat.SharedMedias'}
				component={Components.Chat.SharedMedias}
				options={BackgroundHeaderScreenOptions({
					title: t('chat.shared-medias.title'),
					...CustomTitleStyle(),
					presentation: 'formSheet',
				})}
			/>
			{/* Settings */}
			<NavigationStack.Screen
				name={'Settings.Home'}
				component={Components.Settings.Home}
				options={{
					headerStyle: {
						backgroundColor: colors['secondary-background'],
					},
					headerTintColor: colors['main-text'],
					headerBackTitleVisible: false,
					title: 'Settings',
				}}
			/>
			<NavigationStack.Screen
				name={'Settings.Network'}
				component={Components.Settings.Network}
				options={{
					headerStyle: {
						backgroundColor: colors['secondary-background'],
					},
					headerTintColor: colors['main-text'],
					headerBackTitleVisible: false,
					title: 'Network',
				}}
			/>
			<NavigationStack.Screen
				name={'Settings.Notifications'}
				component={Components.Settings.Notifications}
				options={{
					headerStyle: {
						backgroundColor: colors['secondary-background'],
					},
					title: 'Notifications',
					presentation: 'formSheet',
				}}
			/>
			<NavigationStack.Screen
				name={'Settings.ContactAndConversations'}
				component={Components.Settings.ContactAndConversations}
				options={{
					headerStyle: {
						backgroundColor: colors['secondary-background'],
					},
					title: 'Contact and Conversations',
					presentation: 'formSheet',
				}}
			/>
			<NavigationStack.Screen
				name={'Settings.Appearence'}
				component={Components.Settings.Appearence}
				options={{
					headerStyle: {
						backgroundColor: colors['secondary-background'],
					},
					title: 'Appearence',
					presentation: 'formSheet',
				}}
			/>
			<NavigationStack.Screen
				name={'Settings.ThemeEditor'}
				component={Components.Settings.ThemeEditor}
				options={{
					headerStyle: {
						backgroundColor: colors['alt-secondary-background-header'],
					},
					headerTintColor: colors['reverted-main-text'],
					title: 'ThemeEditor',
					presentation: 'formSheet',
				}}
			/>
			<NavigationStack.Screen
				name={'Settings.DevicesAndBackup'}
				component={Components.Settings.DevicesAndBackup}
				options={{
					headerStyle: {
						backgroundColor: colors['secondary-background'],
					},
					title: 'Devices and Backup',
					presentation: 'formSheet',
				}}
			/>
			<NavigationStack.Screen
				name={'Settings.Security'}
				component={Components.Settings.Security}
				options={{
					headerStyle: {
						backgroundColor: colors['secondary-background'],
					},
					title: 'Security',
					presentation: 'formSheet',
				}}
			/>
			<NavigationStack.Screen
				name={'Settings.Accounts'}
				component={Components.Settings.Accounts}
				options={{
					headerStyle: {
						backgroundColor: colors['secondary-background'],
					},
					title: 'Accounts',
					presentation: 'formSheet',
				}}
			/>
			<NavigationStack.Screen
				name={'Settings.AboutBerty'}
				component={Components.Settings.AboutBerty}
				options={{
					headerStyle: {
						backgroundColor: colors['secondary-background'],
					},
					title: 'About Berty',
					presentation: 'formSheet',
				}}
			/>
			<NavigationStack.Screen
				name={'Settings.MyBertyId'}
				component={Components.Settings.MyBertyId}
				options={BackgroundHeaderScreenOptions({
					title: t('settings.my-berty-ID.title'),
					...CustomTitleStyle(),
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Settings.TermsOfUse'}
				component={Components.Settings.TermsOfUse}
				options={BackgroundHeaderScreenOptions({
					title: 'Terms of use',
					...CustomTitleStyle(),
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Settings.NetworkMap'}
				component={Components.Settings.NetworkMap}
				options={AltBackgroundHeaderScreenOptions({
					title: t('settings.network-map.title'),
					...CustomTitleStyle(),
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Settings.ServicesAuth'}
				component={Components.Settings.ServicesAuth}
				options={BackgroundHeaderScreenOptions({
					title: t('settings.services-auth.title'),
					...CustomTitleStyle(),
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Settings.DeleteAccount'}
				component={Components.Settings.DeleteAccount}
				options={{
					headerShown: false,
					presentation: 'formSheet',
				}}
			/>
			<NavigationStack.Screen
				name={'Settings.DevTools'}
				component={Components.Settings.DevTools}
				options={AltBackgroundHeaderScreenOptions({
					title: t('settings.devtools.title'),
					...CustomTitleStyle(),
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Settings.FakeData'}
				component={Components.Settings.FakeData}
				options={AltBackgroundHeaderScreenOptions({
					title: t('settings.fake-data.title'),
					...CustomTitleStyle(),
					presentation: 'formSheet',
				})}
			/>

			<NavigationStack.Screen
				name={'Settings.SystemInfo'}
				component={Components.Settings.SystemInfo}
				options={AltBackgroundHeaderScreenOptions({
					title: t('settings.system-info.title'),
					...CustomTitleStyle(),
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Settings.AddDevConversations'}
				component={Components.Settings.AddDevConversations}
				options={AltBackgroundHeaderScreenOptions({
					title: t('settings.add-dev-conversations.title'),
					...CustomTitleStyle(),
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Settings.IpfsWebUI'}
				component={Components.Settings.IpfsWebUI}
				options={AltBackgroundHeaderScreenOptions({
					title: t('settings.ipfs-webui.title'),
					...CustomTitleStyle(),
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Settings.DevText'}
				component={Components.Settings.DevText}
				options={AltBackgroundHeaderScreenOptions({
					title: '',
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Settings.BertyServices'}
				component={Components.Settings.BertyServices}
				options={{ headerShown: false, presentation: 'formSheet' }}
			/>
			<NavigationStack.Screen
				name={'Settings.Roadmap'}
				component={Components.Settings.Roadmap}
				options={BackgroundHeaderScreenOptions({
					title: t('settings.roadmap.title'),
					...CustomTitleStyle(),
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Settings.Faq'}
				component={Components.Settings.Faq}
				options={BackgroundHeaderScreenOptions({
					title: t('settings.faq.title'),
					...CustomTitleStyle(),
					presentation: 'formSheet',
				})}
			/>
			<NavigationStack.Screen
				name={'Settings.PrivacyPolicy'}
				component={Components.Settings.PrivacyPolicy}
				options={BackgroundHeaderScreenOptions({
					title: t('settings.privacy-policy.title'),
					...CustomTitleStyle(),
					presentation: 'formSheet',
				})}
			/>
			{/* Modals */}
			<NavigationStack.Screen
				name={'Modals.ManageDeepLink'}
				component={Components.Modals.ManageDeepLink}
				options={{
					presentation: 'containedTransparentModal',
					animation: 'fade',
					headerShown: false,
				}}
			/>
			<NavigationStack.Screen
				name={'Modals.ImageView'}
				component={Components.Modals.ImageView}
				options={{
					presentation: 'containedTransparentModal',
					headerShown: false,
				}}
			/>
			<NavigationStack.Screen
				name={'Modals.EditProfile'}
				component={Components.Modals.EditProfile}
				options={{
					presentation: 'transparentModal',
					headerShown: false,
					animation: 'fade_from_bottom',
				}}
			/>
		</NavigationStack.Navigator>
	)
})
